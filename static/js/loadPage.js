function loadPage(page){ 
	var xmlhttp;
	if (window.XMLHttpRequest) { // code for IE7+, Firefox, Chrome, Opera, Safari
	    xmlhttp = new XMLHttpRequest();
	}
	else { // code for IE6, IE5
	    xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
	}
	xmlhttp.onreadystatechange = function() {
	    if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
	        if(page == 'user'){
				document.getElementById('jumbo_content').innerHTML= xmlhttp.responseText;
			}else if(page == 'map'){
				document.getElementById('jumbo_content').innerHTML= xmlhttp.responseText;
			}
	    }
	}
	if(page == 'user'){
		xmlhttp.open("GET", "/templates/test.html", true);
	}else if(page == 'map'){
		xmlhttp.open("GET", "/templates/mapTest", true);
	}
	
	xmlhttp.send();
}