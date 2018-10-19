var tableRequester = {
	requestParams: {},
    /**
     * @param config
     *    config.rootSelector 父级元素选择器 <br>
     *    config.triggerSelector 触发元素的选择器 默认最好是统一成 css：ct_request <br>
     *    该元素需要有 data-key data-value data-module 这三个属性。 <br>
     *    config.beforeRequest 是请求的前置通知（用于在请求前做一些处理）beforeRequest(element,event)，返回false则不请求
     *    <br>
     */
    bindRequest: function (config) {
        var root = (config.rootSelector) || '';
        var trigger = (config.triggerSelector) || '.maybach_request';
        $(root).delegate(trigger, 'click', function (event) {
            event.preventDefault();
            var element = event.currentTarget;
            var param;
            if ($.isFunction(config.beforeRequest)) {
                param = config.beforeRequest(element, event);
                if (param === false) {
                    event.preventDefault();
                    return false;
                }
            }
            var key = $(element).attr('data-key');
            var value = $(element).attr('data-value');
            tableRequester.request(config, key, value, param);
        });
    },
	request: function (config, key, value, param) {
        this.requestParams[key] = value;
        
        var paramAll = $.extend($(config.requestForm).formToJSON(), this.requestParams);
        paramAll = $.extend(paramAll, param || {});
        
        $(config.tableContainer).bootstrapTable('refresh', {silent: true, query: paramAll});
    }
}