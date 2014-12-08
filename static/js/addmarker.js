var markerPlaced = 0



function updateDrag(location) {
  var input = $('#loc');
  input.val(location);

}

function placeMarker(location) {
  var elem = document.getElementById("stairname");
  if(markerPlaced == 0){
    var marker = new google.maps.Marker({
                position:location,
                map: map,
                draggable:true,
                animation: google.maps.Animation.DROP,
                title: "New Marker",
              icon: "http://maps.google.com/mapfiles/ms/micons/blue.png"
            });
    var input = $('#loc');
    input.val(location);
  }
    
    markerPlaced = 1
    google.maps.event.addListener(marker, 'dragend', function(event){
    updateDrag(marker.getPosition())
    });
    
    
}

function closeSelf(){
    // do something
       alert("conditions satisfied, submiting the form.");
       window.close();
}

function submitForm(form, type){
    // collect the form data while iterating over the inputs
  var data = {};
  for (var i = 0, ii = form.length; i < ii; ++i) {
    var input = form[i];
    if (input.name) {
      data[input.name] = input.value;
    }
  }

  window.fbAsyncInit = function() {
  FB.init({
    appId      : '562407890559656',
    cookie     : true,  // enable cookies to allow the server to access 
                        // the session
    xfbml      : true,  // parse social plugins on this page
    version    : 'v2.1' // use version 2.1
    });
  };
   FB.api('/me', function(response) {
     console.log('Get: ' + response.id);
      data = JSON.stringify(data);
      getUser(response.id, data);

  });
   sendForm(data, type);
}

function sendForm(data, type) {

  var xhr = new XMLHttpRequest();
  if(type == 'marker'){
    xhr.open('POST','http://79.136.28.106:8888/stair' , true);
  }else if(type == 'comment'){
    data.idStair = parseInt(data.idStair);
    xhr.open('POST','http://79.136.28.106:8888/comment' , true);
  }
  xhr.onload = function(e) {};
  if(type!=undefined){
    xhr.send(JSON.stringify(data));
  }
  initialize();
  
}
