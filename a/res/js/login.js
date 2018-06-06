function findGetParameter(parameterName) {
    var result = null,
        tmp = [];
    var items = location.search.substr(1).split("&");
    for (var index = 0; index < items.length; index++) {
        tmp = items[index].split("=");
        if (tmp[0] === parameterName) result = decodeURIComponent(tmp[1]);
    }
    return result;
}

if (findGetParameter('login') === 'false') {
	var loginFailed = document.getElementById('authentication-failed');
	loginFailed.style.display = "block";
}

if (findGetParameter('loggedOut') === 'true') {
	var loggedOut = document.getElementById('logged-out');
	loggedOut.style.display = "block";
}
