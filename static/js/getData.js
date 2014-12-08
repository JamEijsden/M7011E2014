function getLocations(recData, status, stairID)
{
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);
        console.log(obj);
        if(status == 0){
          loadMap(obj);  
        } else if(status == 1){
          //pass along locations from LoadMap and comments for location
          createMarker(recData,obj)
        }
        //document.getElementById("loadhere").innerHTML='<p>'+obj[1].position+'</p>';
      }
    else{
      return "TOMTE";
    }
  };
  if(status == 0){
      xmlHttp.open( "GET", "http://79.136.28.106:8888/stairs", false );   
  }else if(status == 1){
      xmlHttp.open( "GET", "http://79.136.28.106:8888/comment/"+stairID, false );
  }else{
    
  }
  xmlHttp.send( null );
}

function getUser(id, data, action)
{
  var xmlHttp = null;
  console.log(id + ", " + data + ", " + action);
  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);
        if(action == 'form'){
          data = JSON.parse(data);
          data.user = obj.userID;
          sendForm(data);
        } else if(action == 'comment'){
          createComments(data, obj);
        }
        
      }
    else{
      return "ERROR";
    }
  };
  xmlHttp.open( "GET", "http://79.136.28.106:8888/users/"+id, false );
  xmlHttp.send( null );
}

function uploadPhoto(){
  var xmlHttp = null;
  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {      
      }
    else{
      return "ERROR";
    }
  };
  xmlHttp.open( "GET", "http://79.136.28.106:8888/picture/", false );
  xmlHttp.send( null );

}