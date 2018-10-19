(function($){
	$.DataGrid = function(){
		this.settings = $.extend(true, {}, this.defaults);
	};
	
	$.extend($.DataGrid, {
		defaults: {
			
		},
		prototype: {
			/**
			* 新建一个datagrid表格对象
			*/
			generate: function(){
				var _this = this;
				_this.settings.columns = new Array();
				var trObjs = $("#" + _this.settings.id).find("thead").find("tr");
				$(trObjs).each(function(i){
					var childs = $(this).children();
					$(childs).each(function(index){
						var field = $(this).attr("field");
						var colspan = $(this).attr("colspan");
						if(null == colspan || 1 == colspan){
							_this.settings.columns.push(field);
						}
					});
				});
			},
			
			loadData: function(jsonData){
				var _this = this;
				var options = this.parameters?this.parameters:{};
				//empty data area first
				var tbody = $("#" + _this.settings.id).find("tbody");
				if(0 == tbody.length){
					$("#" + _this.settings.id).append("<tbody></tbody>");
					tbody = $("#" + _this.settings.id).find("tbody");
				}
				var strHtml = "";
				for(var index=0;index < options.length; index++){
					var rowData = options[index];
					strHtml += "<tr";
					if(1 == index%2){
						strHtml += " class='table-rows1'";
					}
					strHtml += ">";
					for(var i=0; i< _this.settings.columns.length; i++){
						strHtml += "<td nowrap='nowrap'>" + rowData[_this.settings.columns[i]] + "</td>";
					}
					strHtml += "</tr>";
				}
				$(tbody).html(strHtml);
			}
		}
	});
	
	/**
	* 使用$("#dataTable").datagrid();来初始化datagrid
	* 使用$("#dataTable").datagrid("loadData", jsonData);加载数据
	*/
	$.fn.datagrid = function(method, parameters){
		var datagridId = $(this).attr("id");
		var datagridObj;
		if(datagridMap.contains(datagridId)){
			datagridObj = datagridMap.get(datagridId);
			//使用datagrid中的方法
			datagridObj.parameters = parameters;
			eval("datagridObj." + method + "();");
		}else{
			//初始化datagrid
			datagridObj = new $.DataGrid();
			datagridObj.settings.id = datagridId;
			datagridObj.generate();
			datagridMap.put(datagridId, datagridObj);
		}
	};
})(jQuery);

var datagridMap = new Map();
