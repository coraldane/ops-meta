document.write("<center><span id='loadingDiv' class='loading'>Loading, please wait...</span></center>");

$(document).ready(function(){
	$("#loadingDiv").remove();
	
	$("BODY").delegate("FORM[name='itemForm']", "submit", function(){
		$("#pageIndex", this).val(1);
	});
});

(function($){
	/**
	 * 替代alert框,调用方式可以为$.message("提示内容");
	*/
	$.message = function(text, options){
		var opts = {title:"提示", content: '<span style="font-size:16px;">'+text+'</span>', time: 2};
		if(null != opts){
			opts = $.extend(opts, options);
		}
		var d = dialog(opts);
		d.showModal();
		if(null != options && null != options.time){
			setTimeout(function () {
			    d.close().remove();
			}, options.time * 1000);
		}
		return d;
	};
	$.fn.alert = function(options){
		var me = this;
		var opts = {type:"alert-success", text: "", cancel: true};
		if(null != opts){
			opts = $.extend(opts, options);
		}
		var strArray = new Array();
		strArray.push('<div class="alert '+ opts.type);
		if(true == opts.cancel){
			strArray.push(' alert-dismissable"><button type="button" class="close" data-dismiss="alert" aria-hidden="true">&times;</button>');
		} else {
			strArray.push('">');
		}
		strArray.push(opts.text);
		strArray.push('</div>');
		$(me).html(strArray.join(''));
	};
	$.fn.formToJSON = function(){
	    var o = {};
	    var a = this.serializeArray();
	    $.each(a, function() {
	        if (o[this.name] !== undefined) {
	            if (!o[this.name].push) {
	                o[this.name] = [o[this.name]];
	            }
	            o[this.name].push(this.value || '');
	        } else {
	            o[this.name] = this.value || '';
	        }
	    });
	    return o;
	};
	/**
	* 将FORM表单或者任意容器中的数据内容进行提交
	*/
	$.fn.goAjax = function(parameters){
		var me = this;
		parameters.type = "POST";
		parameters.dataType = null==parameters.dataType?"json":parameters.dataType;
		var semantic = null==parameters.semantic?false:parameters.semantic;
		var checkLogin = null==parameters.checkLogin?false:parameters.checkLogin;
		parameters.error = function(request, textStatus, errorThrown) {
			if(null != errorThrown){
				console.error("post error, status:" + request.status + ",textStatus:" + textStatus + ", error:" + errorThrown);
			}
		};

		if(checkLogin){
			$.ajax({type: "GET",url: "/user/checkLogin",data:{},dataType: "json",
				success: function(response) {
					if(false == response.success){
						var d = dialog({
							title : '登 录',
							url : '/user/login/panel',
							height: 260, 
							onclose: function(){
								if(true == d.data){
									$(me).doAjaxSubmit(parameters, semantic);
								}
							}
						});
						d.showModal();
					} else {
						$(me).doAjaxSubmit(parameters, semantic);
					}
				}
			});
		} else {
			$(me).doAjaxSubmit(parameters, semantic);
		}
	};
	$.fn.doAjaxSubmit = function(parameters, semantic){
		var postContainerSelector = this;
		parameters.data = null==parameters.data?{}:parameters.data;
		if(semantic){
			jsonData = $(postContainerSelector).formToJSON();
			parameters.data = $.extend(parameters.data, jsonData);
		}
		$(postContainerSelector).ajaxSubmit(parameters);
	};
})(jQuery);



function ajaxhtml(obj,href,fn){
	$.getEx(href, {}, function(data){
		$(obj).html(data);
		if(fn){fn();}
	}, "html");
}

function checkLogin(fn){
	$.ajax({type: "GET",url: "user/checkLogin",data:{},dataType: "json",
		success: function(response) {
			if(true == response.success){
				if(fn){fn();}
			} else {
				var d = dialog({
					title : '登 录',
					url : '/user/login/panel',
					height: 260, 
					onclose: function(){
						if(true == d.data){
							if(fn){fn();}
						}
					}
				});
				d.showModal();
			}
		}
	});
}

//ajax请求封装；url:用来包含发送请求的url字符串；getdata:发送到服务器的数据，自动转成字符串；get请求将附加到url后，最好以键值对格式；如：name=john&location=boston   |   {foo:['aa','bb']}  |  {key:value}
$.getEx = function(url, getdata, callback, dataType) {
    $.ajax({
        type: "GET",
        url: url,
        data: getdata,
        dataType: null == dataType?"json":dataType,//解决跨域（跨域访问就是你在一个域环境下，访问另一个域的内容），用户传递一个callback参数给服务端，然后服务端返回数据时将这个callback参数作为行数么来包裹json数据，客户端就可以随意定制自己的函数来返回数据了；
		jsonp:"callback",
		cache: true,
        success: function(data) {//data:请求成功后的回调函数data();服务器根据dataType来返回
            callback(data);
        },
        error: function(request, textStatus, errorThrown) {
        	console.error("get error, status:" + request.status + ",textStatus:" + textStatus + ", error:" + errorThrown);
        }
    });
};
$.postEx = function(url, getdata, callback, dataType) {
    $.ajax({
        type: "POST",
        url: url,
        data: getdata,
		dataType: null == dataType?"json":dataType,//解决跨域（跨域访问就是你在一个域环境下，访问另一个域的内容），用户传递一个callback参数给服务端，然后服务端返回数据时将这个callback参数作为行数么来包裹json数据，客户端就可以随意定制自己的函数来返回数据了；
		jsonp:"callback",
		cache: true,
        success: function(data) {//data:请求成功后的回调函数data();服务器根据dataType来返回
            callback(data);
        },
        error: function(request, textStatus, errorThrown) {
        	console.error("post error, status:" + request.status + ",textStatus:" + textStatus + ", error:" + errorThrown);
        }
    });
};

/**
 * Date Format
 * @param style date format like 'yyyyMMdd'
 */
Date.prototype.format = function(style) {
  var o = {
    "M+" : this.getMonth() + 1, //month
    "d+" : this.getDate(),      //day
    "h+" : this.getHours(),     //hour
    "m+" : this.getMinutes(),   //minute
    "s+" : this.getSeconds(),   //second
    "w+" : "日一二三四五六".charAt(this.getDay()),   //week
    "q+" : Math.floor((this.getMonth() + 3) / 3),  //quarter
    "S"  : this.getMilliseconds() //millisecond
  }
  if(/(y+)/.test(style)) {
	style = style.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
  }
  for(var k in o){
    if(new RegExp("("+ k +")").test(style)){
      style = style.replace(RegExp.$1, RegExp.$1.length == 1 ? o[k] : ("00" + o[k]).substr(("" + o[k]).length));
    }
  }
  return style;
};
