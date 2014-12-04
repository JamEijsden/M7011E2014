
function loadProfile(profile){
	console.log('DAGS ATT LADDA....');
	document.getElementById('fbphoto').src="http://graph.facebook.com/"+profile.id+"/picture?type=large";
	document.getElementById('fbname').innerHTML=profile.name;

 }
