$(document).ready(function(){
	$(window).scroll(function(){
		if(400 < $(window).scrollTop()){
			if(0 == $(".backTop").length){
				$("BODY").append('<div class="backTop" title="BACK TOP"></div>');
				$(".backTop").click(function(){
					$('html,body').animate({scrollTop: '0px'}, 800);
				});
			}
		}else{
			$(".backTop").remove();
		}
	});
});