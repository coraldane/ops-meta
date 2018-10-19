(function($){
	$.FileUpload = function(){
		this.settings = $.extend(true,{},$.FileUpload.defaults);
		this.fileMap = new Map();
		this.fileNameMap = new Map();
	};
	
	$.extend($.FileUpload,{
		defaults:{
			fileIndex :0,
			existsUploadForm: false,
			dwrEnabled:true
		},
		prototype:{
			/**
			 * 初始化文件显示容器
			 * @param handler
			 */
			initContainer: function(handler){
				if(null == this.settings.container){
					var msgShowDiv = document.createElement("DIV");
					msgShowDiv.id = "upload-msg-div";
					msgShowDiv.setAttribute("class", "upload-msg-div");
					$(msgShowDiv).insertAfter(handler);
					this.settings.container = msgShowDiv;
				}
			},
			/** 添加附件
			 * @param obj 按钮对象或者链接元素对象
			 * @param callback 用于处理上传后文件名的函数
			 * 回调函数如下:function(value){$("#input").val(value);}
			 */
			addAttachment: function(obj, callback, options){
				var _this = this;
				if(null != options){
					_this.settings = $.extend(_this.settings, options);
				}
				if(this.settings.dwrEnabled){
					dwr.engine.setActiveReverseAjax(true);
					this.settings.dwrEnabled = false;
				}
				this.callback = callback;
				//判断有无添加附件的表单在页面中
				if(_this.settings.existsUploadForm){
					$(_this.settings.uploadContainer).show();
				}else{
					var position = $(obj).position();
					var strArray = new Array();
					strArray.push('<div id="upload-div" class="uploadDiv" style="position:absolute;');
					strArray.push('left:'+(position.left + $(obj).outerWidth())+'px;top:');
					strArray.push(position.top +'px;">');
					strArray.push('<form id="fileUploadForm" action="upload.mi" method="post" enctype="multipart/form-data">');
					strArray.push('<input type="file" name="fileUploadInput" id="fileUploadInput" ');
					strArray.push((null==_this.settings.accept?'':' accept="'+_this.settings.accept+'"')+'/></form>');
					strArray.push('</div>');
					$(document.body).append(strArray.join(""));
					_this.settings.uploadContainer = $("#upload-div");
					
					this.initContainer(obj);
					_this.settings.existsUploadForm = true;
					
					$("#fileUploadInput").bind("change", function(){
						if("" == $.trim($(this).val())){
							return;
						}
						var fileName = _this.getFileName($(this).val());
						if("" != $(this).attr("accept")){
							var fileExtension = _this.getFileExtension(fileName);
							if("" == fileExtension){return;}
							var bAccept = false;
							var extArray = $(this).attr("accept").split("|");
							for(var index=0; index < extArray.length; index+=1){
								if(fileExtension == extArray[index]){bAccept=true;break;}
							}
							if(false == bAccept){alert("允许输入的文件后缀名为:"+_this.settings.accept);return;}
						}
												
						_this.settings.fileIndex ++;
						_this.fileNameMap.put(_this.settings.fileIndex, fileName);
						FileUploadListener.startUpload(_this.settings.fileIndex);
						$($(this).parent()).ajaxSubmit();
						$(_this.settings.uploadContainer).hide();
					});
				}
			},
			updateProgress: function(fileIndex, uploadInfo){
				var _this = this;
				var fileName = _this.fileNameMap.get(fileIndex);
				var totalSize = uploadInfo.totalFileSize;
				var percent = parseInt(uploadInfo.percent);
				var attachmentId = uploadInfo.attachmentId;
				if(0 < attachmentId){
					_this.fileMap.put(fileIndex, attachmentId);
				}
				if($("#processbar"+fileIndex).size() == 0){
					var str = '<div class="process-info">' + fileName +
					'&nbsp;&nbsp;<span class="process-bar"><span id="processbar'+fileIndex+
					'" class="process-bar-inner"></span></span>&nbsp;&nbsp;<span id="process-bar-percent'+fileIndex +
					'">0%</span>&nbsp;&nbsp;&nbsp;<a href="javascript:void(0);" onclick="fileUpload.cancelUpload('+
					fileIndex +')">删除</a>&nbsp;&nbsp;&nbsp;' + totalSize +'</div>';
					$(_this.settings.container).append(str);
								
					//call callback function
					_this.callback(_this.getAttachments());
				}
				$("#processbar"+fileIndex).width(percent);
				$("#process-bar-percent"+fileIndex).html(percent+"%");
			},
			/**返回已经上传的附件编码列表,对应sys_attachment表中的attachment_id */
			getAttachments: function(){
				var attachmentIds = "";
				var keys = this.fileMap.keySet();
				for(var i=0; i< keys.length; i++){
					var key = keys[i];
					attachmentIds += this.fileMap.get(key);
					if(i < keys.length -1){
						attachmentIds += ",";
					}
				}
				return attachmentIds;
			},
			getFileExtension: function(fileName){
				var start = fileName.lastIndexOf(".");
				return (0>start?"":fileName.substring(start+1));
			},
			getFileName: function(filepath){
				var str = filepath.replace(/\\/g, "/");
				var end = str.lastIndexOf("/");
				return str.substring(end+1);
			},
			cancelUpload: function(fileIndex){
				var _this = this;
				FileUploadListener.cancelUpload(fileIndex, {
					callback:function(){
						_this.fileNameMap.remove(fileIndex);
						_this.fileMap.remove(fileIndex);
						var processBar = $("#processbar"+fileIndex);
						$(processBar).parent().parent().remove();
					}
				});
			},
			/**
			* 根据附件文件名称所在的文本框初始化附件列表
			* @param obj
			*/
			init: function(obj){
				var _this = this;
				var fileNames = obj.value;
				if("" == $.trim(fileNames)){
					return;
				}
				this.initContainer(obj);
				$.ajax({
					type: "POST",
					url: "upload.mi",
					data:{method:'getAttachmentNames', ids: fileNames},
					dataType: "json",
					success: function(data) {
						var index = 0;
						for(index =0; index < data.length; index ++){
							var id = index +1;
							var info = data[index];
							_this.fileMap.put(id, info.attachmentId);
							_this.fileNameMap.put(id, info.displayName);
							var str = '<div class="process-info">' + info.displayName +
							'&nbsp;&nbsp;<a href="javascript:void(0);" onclick="fileUpload.deleteFile(this,'+
							info.attachmentId +')">删除</a>&nbsp;&nbsp;&nbsp;' + info.fileSizeDesc +'</div>';
							$(_this.settings.container).append(str);
						}
						_this.settings.fileIndex = index +1;
					}
				});
			},
			deleteFile: function(handler, id){
				var _this = this;
				$.ajax({
					type: "POST",
					url: "upload.mi",
					data:{method:'deleteUploadFile', attachmentId: id},
					dataType: "html",
					success:function(){
						_this.fileMap.remove(id);
						_this.fileNameMap.remove(id);
						$(handler).parent().remove();
					}
				});
			}
		}
	});
})(jQuery);

/**创建一个FileUpload对象*/
var fileUpload = new $.FileUpload();

function updateFileUploadProgress(fileIndex, uploadInfo){
	fileUpload.updateProgress(fileIndex, uploadInfo);
}

function callbackFileUpload(fileIndex, attachId){
	fileUpload.fileMap.put(fileIndex, attachId);
}