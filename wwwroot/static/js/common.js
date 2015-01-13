// 获取 access_token, url 格式一般为 http://aaa.bbb.com/ccc?xxx=yyy&access_token=tttttt
function getAccessToken() {
	var urlQueryString = window.location.search;
	var tokenIndex = urlQueryString.indexOf("?access_token=");
	if (tokenIndex === -1) {
		tokenIndex = urlQueryString.indexOf("&access_token=");
		if (tokenIndex === -1) {
			return "";
		}
	}
	urlQueryString = urlQueryString.slice(tokenIndex+14);
	var endIndex = urlQueryString.indexOf("&");
	if (endIndex === -1) {
		endIndex = urlQueryString.indexOf("#");
		if (endIndex === -1) {
			endIndex = urlQueryString.length;
		}
	}
	return decodeURIComponent(urlQueryString.slice(0, endIndex));
}

// 管理员登出事件监听器, 需要在 JQuery 之后加载
function adminLogoutEventListener(event) {
	event.preventDefault();
	
	$.ajax({
		"url": "/admin/logout?access_token=" + event.data.access_token,
		"type": "GET",
		"dataType": "json",
		"error": function(jqXHR, textStatus, errorThrown) {
			window.alert("请求出现错误");
		},
		"success": function(data, textStatus, jqXHR){
			switch (data.err_code) {
			case 0:
				window.alert("登出成功");
				window.location.href = "/admin/login.html";
				return;
			case 401:
				window.alert("您没有登录, 请先登录");
				window.location.href = "/admin/login.html";
				return;
			default:
				window.alert(data.err_msg);
				return;
			}
		}
	});
}

// 商家登出事件监听器, 需要在 JQuery 之后加载
function merchantLogoutEventListener(event) {
	event.preventDefault();
	
	$.ajax({
		"url": "/merchant/logout?access_token=" + event.data.access_token,
		"type": "GET",
		"dataType": "json",
		"error": function(jqXHR, textStatus, errorThrown) {
			window.alert("请求出现错误");
		},
		"success": function(data, textStatus, jqXHR){
			switch (data.err_code) {
			case 0:
				window.alert("登出成功");
				window.location.href = "/merchant/login.html";
				return;
			case 401:
				window.alert("您没有登录, 请先登录");
				window.location.href = "/merchant/login.html";
				return;
			default:
				window.alert(data.err_msg);
				return;
			}
		}
	});
}