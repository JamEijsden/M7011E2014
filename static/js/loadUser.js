function loadUser(data){
	console.log('TOMTE');
	document.getElementById('userTitle').innerHTML = data.first_name + " " + data.last_name;
	getStairUser(data.userID);
	
}

function createUserLocations(data){
	parent = document.getElementById('userLoc');
	if(data == null){
		parent.innerHTML = 'No locations to show';
		return;
	}
	
	$('userLoc').empty();
	for(var i = 0; i < data.length;i++){
		var loc = document.createElement('input');
		loc.id = data[i].id;
		loc.type = 'button';
		loc.className = "btn btn-primary";
		loc.name = data[i].user;
		loc.value = data[i].stairname;
		console.log(loc);
		parent.appendChild(loc);
		
		loc.onclick = function(ev){
			var elem =ev.target;
			openModal(elem.id, elem.name);
			
		};
	}
}

function openModal(id, name){
		 console.log(id + " " + name);
    	 getStair(id,'' ,'user');
   		 getPreviewStair(id);
   		 document.getElementById('modalUser').value = name;
   		 document.getElementById('modalStair').value = id;	
}