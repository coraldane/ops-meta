/**
 * Created by coral on 08/05/2017.
 */
/**
 * Created by lxr on 2017/3/23.
 */

$(function(){
    $('.tab-menu>a').click(function(){
        var $obj = $(this),_index = $obj.index();
        !$obj.hasClass('active') ? $obj.addClass('active').siblings().removeClass('active') : '';
        $('.pane-list>.tab-pane').eq(_index).addClass('active').siblings().removeClass('active');
    });
    var str = '<div id="weixin-tip" class="download"><img src="/static/images/global/android.png" alt="微信扫描打开APP下载链接提示代码优化" /></div>'
        + '<div id="ios-weixin-tip" class="download"><img src="/static/images/global/ios.png" alt="微信扫描打开APP下载链接提示代码优化" /></div>';
    $('body').append(str);
});

/*判断是否在微信*/
var ua = navigator.userAgent.toLowerCase();
var is_weixin = function () {
    if (navigator.userAgent.match(/MicroMessenger/i)) {
        return true;
    } else {
        return false;
    }
}

function isPhone(id){
    var tip = $("#"+id);
    tip.show();
    tip.click(function () {
        tip.hide();
    });
}

function openApp(openNew, roomId) {
	if(null != roomId && "" != roomId){
		openNew = true
	}
	if (is_weixin()) {
        if(navigator.userAgent.match(/(iPhone|iPod|iPad)/i)){
            isPhone("ios-weixin-tip");
        } else if(navigator.userAgent.match(/android/i)){
            isPhone("weixin-tip");
        }
    } else {
		if (navigator.userAgent.match(/(iPhone|iPod|iPad)/i)) {
			var ifr = document.createElement("iframe");
			ifr.src = targetURLIOS;
			ifr.style.display = "none";
			document.body.appendChild(ifr);
			
			window.setTimeout(function () {
				document.body.removeChild(ifr);
				window.location.href = downloadURLIOS;
			}, 2000);
			
			if(openNew){
				window.location.href = targetURLIOS + (null==roomId?"":roomId);
			}
		} else if (navigator.userAgent.match(/android/i)) {
			var loadTime = +(new Date());
			window.setTimeout(function () {
				var timeOut = +(new Date());
				if (timeOut - loadTime < 5000)
				{
					window.location.href = downloadURLAndroid;

					hasApp = false;
				} else {
					window.close();
				}
			}, 1000);

			if(openNew){
				window.location.href = targetURLAndroid + (null==roomId?"":roomId);
			}
		}
	}
}