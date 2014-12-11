function loadUser(data){
	console.log('TOMTE');
	document.getElementById('userTitle').innerHTML = data.first_name + " " + data.last_name;
	
}

function createUserLocations(){
	
	
}

function openModal(){
	$('#modal2').on('show.bs.modal', function () {
	  $('.modal-content').css('height',515);
	  //$('.modal-content').css('max-height','500px');
	});
	$('#modal2').modal('show');
	('#modal-body').css('height',$('#modal-content').height()*0.50);	
}