var _validate_messages = {
    required: "必选字段",
    /*remote: "请修正该字段",*/
    email: "电子邮件格式错误",
    domain: "请输入正确的域名",
    url: "非法网址",
    date: "非法日期",
    dateISO: "日期格式错误 (ISO).",
    number: "非法数字",
    digits: "只能输入整数",
    mobile: "格式错误",
    creditcard: "请输入合法的信用卡号",
    equalTo: "请再次输入相同的值",
    accept: jQuery.format("允许输入的文件后缀名为 {0}"),
    maxlength: jQuery.format("最大长度为 {0} "),
    minlength: jQuery.format("最小长度为 {0}"),
    rangelength: jQuery.format("长度范围为:{0}-{1}"),
    range: jQuery.format("有效值为 {0}-{1} 之间"),
    max: jQuery.format("最大值为 {0}"),
    min: jQuery.format("最小值为 {0}"),
    byterangelength: jQuery.format("长度范围为:{0}-{1}(每个中文占2个字节)"),
    scale: jQuery.format("小数点后最多{0}位")
};

(function($){
	$.fn.addValidator = function(){
		var $form = this;
		$("INPUT[type!='hidden'], TEXTAREA, SELECT", $form).each(function(index){
			var el = this;
			var ruleAtts = new Object();
			$.each(el.attributes, function(index, attr){
				var attrName = attr.name;
				if("required" == attrName){
					ruleAtts[attrName] = true;
				}else if("remote" == attrName){
					var data = eval("document." + attr.value);
					ruleAtts[attrName] = data;
				}else if(null != _validate_messages[attrName]){
					if(attrName.toLowerCase().indexOf("range") >= 0){
						ruleAtts[attrName] = eval(attr.value);
					}else{
						ruleAtts[attrName] = attr.value;
					}
				}else if("validateChar" == attrName){
					ruleAtts[attrName] = true;
				}
			});
			if("{}" != JSON.stringify(ruleAtts)){
				$(el).rules("add", ruleAtts);
			}
		});
	};

	$.fn.removeValidator = function(){
		var $form = this;
		$("INPUT[type!='hidden'], TEXTAREA, SELECT", $form).each(function(index){
			var el = this;
			var ruleAtts = new Array();
			$.each(el.attributes, function(index, attr){
				var attrName = attr.name;
				if("validateChar" == attrName || null != _validate_messages[attrName]){
					ruleAtts.push(attrName);
				}
			});
			if(0 < ruleAtts.length){
				$(this).rules("remove", ruleAtts.join(" "));
			}
		});
	};
})(jQuery);

jQuery.extend(jQuery.validator.messages, _validate_messages);

$(document).ready(function(){
	/***********************用于处理前端验证规则 Start********************************/
	$.validator.setDefaults({
		submitHandler: function(form){form.submit();}
	});

	jQuery.validator.addMethod("domain", function(value, element) {
		  return this.optional(element) || /^[0-9a-zA-Z]+[0-9a-zA-Z\.-]*\.(com(.cn)?|cn|com.hk|info|me|in|io|net(.cn)?|org(.cn)?|gov(.cn)?|edu(.cn)?|sh|hk|cc|co|tv|asia|pw|top|biz|mobi|tm|tw|host|us|name)$/i.test(value);
	}, _validate_messages.domain);
	
	jQuery.validator.addMethod("mobile", function(value, element) {
		  return this.optional(element) || /^0?1(([3578][0-9]{1})|(59)){1}[0-9]{8}$/i.test(value);
	}, _validate_messages.domain);

	// 中文字两个字节
	jQuery.validator.addMethod("byterangelength", function(value, element, param) {
		var length = value.length;
		for(var i = 0; i < value.length; i+=1){
			if(value.charCodeAt(i) > 127){
				length+=1;
			}
		}
	  return this.optional(element) || ( length >= param[0] && length <= param[1] );
	}, _validate_messages.byterangelength);
	
	// 小数点后多少位
	jQuery.validator.addMethod("scale", function(value, element, param) {
		var start = value.indexOf(".");
		var length = 0;
		if(0 < start){
			for(var i = start+1; i < value.length; i+=1){
				length += 1;
			}
		}
	  return this.optional(element) || (0 == start || length <= param);
	}, _validate_messages.scale);

	//有效字符验证
	jQuery.validator.addMethod("validateChar", function(value, element) {
	  return this.optional(element) || /^[\u003a-\uFFE5\w]+$/.test(value);
	}, "只能包括中英文、数字和下划线");//0391 - A

	//密码字符验证
	jQuery.validator.addMethod("password", function(value, element) {
	  return this.optional(element) || /^[a-zA-Z0-9_@]*$/.test(value);
	}, "只能包括英文字母、数字和下划线");

	$.fn.getAttributes = function() {
		var attributes = new Map();
		if(!this.length)
		return this;
		$.each(this[0].attributes, function(index, attr) {
			attributes.put(attr.name, attr.value);
		});
		return attributes;
	}

	$("FORM.validate").each(function(i){
		var $form = this;
		var rules = new Array();
		var msgs = new Array();
		$("INPUT[type!='hidden'], TEXTAREA, SELECT", $form).each(function(index){
			var el = this;
			var name = $(el).attr("name");
			var ruleAtts = new Array();
			var msgAtts = new Array();
			$.each(el.attributes, function(index, attr){
				var attrName = attr.name;
				if("required" == attrName){
					ruleAtts[attrName] = true;
					msgAtts[attrName] = attr.value;
				}else if("remote" == attrName){
					var data = eval("document." + attr.value);
					ruleAtts[attrName] = data;
					msgAtts[attrName] = data.msg;
				}else if(null != _validate_messages[attrName]){
					if(attrName.toLowerCase().indexOf("range") >= 0){
						ruleAtts[attrName] = eval(attr.value);
					}else{
						ruleAtts[attrName] = attr.value;
					}
					msgAtts[attrName] = _validate_messages[attrName];
				}else if("validateChar" == attrName){
					ruleAtts[attrName] = true;
					msgAtts[attrName] = "只能包括中英文、数字和下划线";
				}
			});
			rules[name] = ruleAtts;
			msgs[name] = msgAtts;
		});
		var errorPlacement = $($form).attr("errorPlacement");
		var fnError;
		if(undefined != errorPlacement){
			fnError = eval("document." + errorPlacement);
		}else{
			fnError = function(error, element) {
				error.appendTo( element.next());
			}
		}

		$($form).validate({
			rules: rules,
			messages: msgs,
			/* 错误信息的显示位置 */
			errorPlacement: fnError,
			/* 验证通过时的处理 */
			success: function(label){
				label.html(" ");
			},
			/* 获得焦点时不验证 */
			focusInvalidate: true,
			onkeyup: false
		});
	});
	/***********************用于处理前端验证规则 End*********************************/
});