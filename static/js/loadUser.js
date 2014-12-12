function loadUser(data){
	console.log('TOMTE');
	document.getElementById('userTitle').innerHTML = data.first_name + " " + data.last_name;
	
}

function createUserLocations(data){
	parent = document.getElementById('userLoc');
	if(){
		parent.innerHTML = 'No locations to show';
		return;
	}
	$('userLoc').empty();
	for(var i = 0; i < data.length;i++){
		var loc = document.createElement('input');
		loc.id = data.id;
		loc.class = "btn btn-primary";
		loc.onclick = "openModal();";
		loc.value = data.stairname;
		parent.appendChild(loc);
	}
	
	
}

function openModal(){
	$('#modal2').on('show.bs.modal', function () {
	  $('.modal-content').css('height',515);
	  //$('.modal-content').css('max-height','500px');
	});
	$('#modal2').modal('show');
	('#modal-body').css('height',$('#modal-content').height()*0.50);	
}