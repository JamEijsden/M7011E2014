function getLocations(recData, status, stairID)
{
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);
        console.log(obj);
        loadMap(obj);  
        
    }else{
      return "TOMTE";
    }
  };
  xmlHttp.open( "GET", "http://79.136.28.106:8888/stairs", false );   
  xmlHttp.send( null );
}

function getUserStairs(user_id){
   var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);  
    }else{
      return "TOMTE";
    }
  };
  
     xmlHttp.open( "GET", "http://79.136.28.106:8888/stairs/"+user_id, false );
  
  xmlHttp.send( null ); 

}

function getStair(id, marker){
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);  
      if(marker.photo == ''){
        console.log('Finish loading pic');
        marker.photo = obj.photo;
        getStair(marker.id, marker);
      }else{        
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json); 
        console.log('Finish loading stair: ' + obj); 
        appendToMarker(obj, marker);
      }
    }else{
      return "TOMTE";
    }
  };
  if(marker.photo == ''){
    console.log('Begun loading pic');
    xmlHttp.open( "GET", "http://79.136.28.106:8888/stair/photo/"+id, false );   
  }else{
    console.log('Loading stair only');
     xmlHttp.open( "GET", "http://79.136.28.106:8888/stair/"+id, false );
  }
  xmlHttp.send( null ); 
}


function getComments(stairID){
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);
        getUserFromComment(obj);  
        
    }else{
      return "Error";
    }
  };
  xmlHttp.open( "GET", "http://79.136.28.106:8888/comment/"+stairID, false );   
  xmlHttp.send( null );
}

function postComment(form){
  document.getElementById('modalComment').value = '';
  var data = {};
  for (var i = 0, ii = form.length; i < ii; ++i) {
    var input = form[i];
    if (input.name) {
      data[input.name] = input.value;
    }
  }
  data.idStair = parseInt(data.idStair);
  console.log(data);
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        console.log("Comment successfully posted");
        getComments(data.idStair);
    }else{
      return "Error";
    }
  };
  xmlHttp.open( "POST", "http://79.136.28.106:8888/comment", true );   
  xmlHttp.send(JSON.stringify(data));

}

function getUser(id, data, action){
  var xmlHttp = null;
  console.log(id + " " + data + " " + action); 
  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);

        if(action == 'form'){
          data = JSON.parse(data);
          data.user = obj.userID;
          sendForm(data, action);

        } else if(action == 'comment'){
          createComments(data, obj);
          
        }else if(action == 'user'){
          loadUser(obj);
        }else if('me'){
            $('me').value = obj.userID;
        }
        
      }
    else{
      return "ERROR";
    }
  };
  xmlHttp.open( "GET", "http://79.136.28.106:8888/users/"+id, false );
  xmlHttp.send( null );
}

function uploadPhoto(form){


  var data = {};
  for (var i = 0, ii = form.length; i < ii; ++i) {
    var input = form[i];
    if (input.name) {
      data[input.name] = input.value;
    }
  }
  data.userID = parseInt(data.userID);
  data.idStair = parseInt(data.idStair);
  var photo = document.getElementById('modalImage');
  if(photo.files.length){
    var reader = new FileReader();
        function success(evt){
          data.photo = evt.target.result; 
          send(data);
          //alert(evt.target.result);

        };
        reader.onload = success;
        
     reader.readAsDataURL(photo.files[0]);
       
         
      
          
        
         
/*
         var xmlHttp = null;
         xmlHttp = new XMLHttpRequest();
         xmlHttp.onreadystatechange=function() {
          if (xmlHttp.readyState==4 && xmlHttp.status==200) { 
          console.log('UPLOAD SUCCESS');     
          }else{
            return "ERROR";
          }
        };
        xmlHttp.open( "POST", "http://79.136.28.106:8888/picture", true );
        xmlHttp.send(JSON.stringify(data)); 
  */  }
  }
 
function send(data){
          
          
          console.log(data);

          xmlHttp = new XMLHttpRequest();
         xmlHttp.onreadystatechange=function() {
          if (xmlHttp.readyState==4 && xmlHttp.status==200) {

            getPreviewStair(data.idStair);
            var res = JSON.parse(xmlHttp.responseText); 
            console.log(res);
            console.log('UPLOAD SUCCESS'); 

            var suc = document.getElementById('suc');
            suc.innerHTML =  "Upload successful!";
            setTimeout(function(){console.log('TIMOUT DONE');document.getElementById('suc').innerHTML ="Upload picture";},3000);   
            }else{
            return "ERROR";
          }
        };
        xmlHttp.open( "POST", "http://79.136.28.106:8888/picture", true );
        xmlHttp.send(JSON.stringify(data));
}


function getPreviewStair(stairID){
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
   xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);
        createPhotos(obj);  
        
    }else{
      return "Error";
    }
  };
  xmlHttp.open( "GET", "http://79.136.28.106:8888/stair/picture/preview/"+stairID, false );   
  xmlHttp.send( null );
}


function getPreviewUser(userID){
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);
        createPhotos(obj);  
        
    }else{
      return "Error";
    }
  };
  xmlHttp.open( "GET", "http://79.136.28.106:8888/users/picture/preview/"+userID, false );   
  xmlHttp.send( null );
}

