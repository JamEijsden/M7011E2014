var markerPlaced = 0
 var marker;


function updateMarker(location) {
  var input = $('#loc');
  input.val(location);

}

function placeMarker(location) {
  var elem = document.getElementById("stairname");
 
  if(markerPlaced == 0){
        marker = new google.maps.Marker({
                position:location,
                map: map,
                draggable:true,
                animation: google.maps.Animation.DROP,
                title: ""+location,
              icon: "http://maps.google.com/mapfiles/ms/micons/blue.png"
            });
        updateMarker(marker.getPosition());
  } else{
      console.log(marker.getTitle() + " " + marker.getPosition() + " " + location);
      marker.setPosition(location);
      updateMarker(marker.getPosition());
    return;
  }

    
    markerPlaced = 1
    google.maps.event.addListener(marker, 'drag', function(event){
       updateMarker(marker.getPosition());
    });

    google.maps.event.addListener(marker, 'dragend', function(event){
        updateMarker(marker.getPosition());
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
  var image = document.getElementById('image').files;
  if(image.length){
    var reader = new FileReader();
        function success(evt){
          data.photo = evt.target.result;
          prepareForm(data, type);
          //alert(evt.target.result);

        }

        reader.onload = success;
        
     reader.readAsDataURL(image[0]);

  
   //sendForm(data,'form');
}
}

function prepareForm(data, type){
  
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
      data = JSON.stringify(data);
      getUser(response.id, data, type);

  });

}

function sendForm(data, type) {

  var xhr = new XMLHttpRequest();

    data.user = parseInt(data.user);
    xhr.open('POST','http://79.136.28.106:8888/stair' , true);
  
  xhr.onreadystatechange=function() {
    if (xhr.readyState==4 && xhr.status==200) {
      console.log("SUCCESSFULLY UPLOADED");
    }
  }
  if(type!=undefined){
    xhr.send(JSON.stringify(data));
  }
  document.getElementById('newLoc').reset();
  initialize();
  markerPlaced = 0;
  return false;
  
}
