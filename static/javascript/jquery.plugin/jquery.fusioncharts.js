/**
//Multi-series Line 2D example
$("#chartDiv").FusionChart({chartType:"MSLine", caption:"日销量趋势",
dataList: jsonData.dataList, width:820, height:200,
xValues:[{name:"销售日期", key:"trans_time"}], 
yValues:[{name:"即时到账", key:"trans_value1"},
	{name:"担保交易", key:"trans_value2"},
	{name:"预存款支付", key:"trans_value3"}]
});

//Pie3D example
$("#chartDiv").FusionChart({chartType:"Pie3D", caption:"日销量趋势",
dataList: jsonData.dataList, width:820, height:200,
xValues:[{name:"销售日期", key:"create_time"}], 
yValues:[{name:"交易成功", key:"tradeNum4"}],
showPercentageValues:"1"
});

//Column 3D + Line Single Y example
$("#chartDiv").FusionChart({chartType:"MSColumnLine3D", caption:"日销量趋势",
dataList: jsonData.dataList, width:820, height:200,
xValues:[{name:"销售日期", key:"create_time"}], 
yValues:[{name:"等待付款", key:"tradeNum1"},
		 {name:"等待发货", key:"tradeNum2"},
		 {name:"发货中", key:"tradeNum3"},
		 {name:"交易成功", key:"tradeNum4", renderAs:"Line"},
		 {name:"交易关闭", key:"tradeNum5"},
		 {name:"退款中", key:"tradeNum6"}]
});
*/
(function ($) {
	$.FusionChart = function(){
		this.settings = $.extend(true, {}, $.FusionChart.defaults);
		this.FLASH_FILE_PATH = $("base").attr("href") +"javascript/FusionCharts/";
		this.colorArray=[
			"1941A5", //Dark Blue
			"CCCC00", //Chrome Yellow+Green
			"999999", //Grey
			"0099CC", //Blue Shade
			"FF0000", //Bright Red 
			"006F00", //Dark Green
			"0099FF", //Blue (Light)
			"FF66CC", //Dark Pink
			"669966", //Dirty green
			"7C7CB4", //Violet shade of blue
			"FF9933", //Orange
			"9900FF", //Violet
			"99FFCC",//Blue+Green Light
			"CCCCFF", //Light violet
			"669900", //Shade of green
			"AFD8F8",
			"F6BD0F",
			"8BBA00",
			"A66EDD",
			"F984A1"];
	};
	
	$.extend($.FusionChart, {
		/**
		xAxisName:"",//X轴名称
		yAxisName:"",//Y轴名称
		*/
		defaults: {
			chartType:"Column2D",
			chartId: "columnChart",
			caption: "Column2D Chart Demo",
			width: 400,
			height: 300,
			dataList: [],//for example jsonData.dataList
			xValues:[{name:"销售日期", key:"trans_date"}],
			yValues:[{name:"即时到账", key:"tran_value1"},{name:"担保交易", key:"tran_value2"}],
			showValues: "0",
			showBorder:"0",//Whether to show a border around the chart or not?
			bgColor:"FFFFFF",
			bgAlpha:"40,100",
			chartLeftMargin: 15,
			chartRightMargin: 30,
			chartTopMargin: 5,
			chartBottomMargin: 5,
			drawAnchors: "0" //Whether to draw anchors on the chart? If the anchors are not shown, then the tool tip and links won't work.
		},
		prototype: {
			render: function(options){
				$.extend(true, this.settings, options);
				var chart = new FusionCharts(this.FLASH_FILE_PATH+options.chartType+'.swf', options.chartId+"_1", this.settings.width, this.settings.height);
				eval("var xmlArray = this.render"+options.chartType+"();");
				//alert(xmlArray.join(""));
				chart.setDataXML(xmlArray.join(""));
				chart.render(options.chartId);
			},
			/**
			* 初始化FusionChart XML Header
			*/
			generateHeader: function(){
				var optArgs = ["xAxisName","yAxisName","drawAnchors","showPercentageValues","numberPrefix",
					"decimalPrecision"];
				var strArray = new Array();
				var options = this.settings;
				strArray.push("<chart caption='"+options.caption+"'");
				for(var i=0; i<optArgs.length; i+=1){
					if(null!=options[optArgs[i]]){strArray.push(" "+optArgs[i]+"='"+options[optArgs[i]]+"'");}
				}
				strArray.push(" showBorder='"+options.showBorder+"' canvasBorderThickness='1' canvasBorderAlpha='20'");
				strArray.push(" bgColor='"+options.bgColor+"' bgAlpha='"+options.bgAlpha+"'");
				strArray.push(" chartLeftMargin='"+options.chartLeftMargin+"' chartRightMargin='"+options.chartRightMargin+"'");
				strArray.push(" chartTopMargin='"+options.chartTopMargin+"' chartBottomMargin='"+options.chartBottomMargin+"'");
				strArray.push(" showValues='"+options.showValues+"'>");
				return strArray;
			},
			/** Pie 2D */
			renderPie2D: function(){
				return this.renderPie3D();
			},
			/** Pie 3D */
			renderPie3D: function(){
				var options = this.settings;
				var strArray = this.generateHeader();
				//Start to parse xValues[0] and yValues[0]
				for(var i=0; i< options.dataList.length; i+=1){
					var element = options.dataList[i];
					strArray.push("<set label='"+element[options.xValues[0].key]+"' value='"+element[options.yValues[0].key]+"'/>");
				}
				strArray.push("</chart>");
				return strArray;
			},
			/** Area 2D */
			renderArea2D: function(){
				return this.renderColumn3D();
			},
			/** Column 2D */
			renderColumn2D: function(){
				return this.renderColumn3D();
			},
			/** Column 3D */
			renderColumn3D: function(){
				var options = this.settings;
				var strArray = this.generateHeader();
				//Start to parse xValues[0] and yValues[0]
				for(var i=0; i< options.dataList.length; i+=1){
					var element = options.dataList[i];
					strArray.push("<set label='"+element[options.xValues[0].key]+"' value='"+element[options.yValues[0].key]+"'/>");
				}
				//Start to parse trade lines.
				if(null != options.trendlines && 0 < options.tradelines.length){
					strArray.push("<tradelines>");
					for(i=0; i<options.tradelines.length; i+=1){
						var line = options.tradelines[i];
						strArray.push("<line startValue='"+line.startValue+"' displayValue='"+line.displayValue+"'");
						if(null != line.color){strArray.push(" color='"+line.color+"'");}
						if(null != line.showOnTop){strArray.push(" showOnTop='"+line.showOnTop+"'");}
						strArray.push("/>");
					}
					strArray.push("</tradelines>");
				}
				strArray.push("</chart>");
				return strArray;
			},
			/** Multi-series Column 2D  */
			renderMSColumn2D: function(){
				return this.renderMSLine();
			},
			/** Multi-series Column 3D */
			renderMSColumn3D: function(){
				return this.renderMSLine();
			},
			/** Multi-series Area 2D */
			renderMSArea: function(){
				return this.renderMSLine();
			},
			/** Multi-series Line 2D */
			renderMSLine: function(){
				var options = this.settings;
				var strArray = this.generateHeader();
				//Start to parse xValues data.
				strArray.push("<categories>");
				for(var i=0; i<options.dataList.length; i+=1){
					strArray.push("<category label='"+options.dataList[i][options.xValues[0].key]+"'/>");
				}
				strArray.push("</categories>");
				
				//Start to parse yValues data.
				for(i=0; i<options.yValues.length; i+=1){
					var element = options.yValues[i];
					strArray.push("<dataset seriesName='"+element.name+"'");
					if(null!=element["renderAs"]){strArray.push(" renderAs='"+element.renderAs+"'");}
					strArray.push(">");
					for(var j=0; j<options.dataList.length; j+=1){
						strArray.push("<set value='"+options.dataList[j][element.key]+"'/>");
					}
					strArray.push("</dataset>");
				}
				
				//Start to parse trade lines.
				if(null != options.trendlines && 0 < options.tradelines.length){
					strArray.push("<tradelines>");
					for(i=0; i<options.tradelines.length; i+=1){
						var line = options.tradelines[i];
						strArray.push("<line startValue='"+line.startValue+"' displayValue='"+line.displayValue+"'");
						if(null != line.color){strArray.push(" color='"+line.color+"'");}
						if(null != line.showOnTop){strArray.push(" showOnTop='"+line.showOnTop+"'");}
						strArray.push("/>");
					}
					strArray.push("</tradelines>");
				}
				strArray.push("</chart>");
				return strArray;
			},
			/** Column 3D Line (Single Y) Combination Chart Specification Sheet */
			renderMSColumnLine3D: function(){
				return this.renderMSLine();
			}
		}
	});
	
	$.fn.FusionChart = function(options){
		var errorMsg = "";
		var methodName = "render" + options.chartType;
		if(false == $.FusionChart.prototype.hasOwnProperty(methodName)){
			//errorMsg = "This jquery plugin doesn't support the chart type at all.";
			errorMsg = "<div style='margin-top:10px;text-align:center;'><font color='red'>这个jQuery插件不支持在所有的图表类型！</font></div>";
		}else if(null == options.dataList || 0 == options.dataList.length){
			//errorMsg = "Argument [dataList] is required and must be not empty.";
			errorMsg = "<div style='margin-top:10px;text-align:center;'><font color='red'>无法显示图表，数据是必需的，必须是不为空！-_-</font></div>";
		}
		
		if("" != errorMsg){
			$(this).html(errorMsg);
			return false;
		}
		var parameters = {chartId: $(this).attr("id")};
		$.extend(true, parameters, options);
		new $.FusionChart().render(parameters);
	};
})(jQuery);