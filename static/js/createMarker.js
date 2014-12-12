var currentMark
var loc;
function createMarker(location) {
  var marker = new google.maps.Marker({
      id: location.id,
      position: location.position,
      map: map,
      title: location.stairname,
      description: location.description,
      photo:"",
      creator:location.user

  });   
  document.getElementById('myModalLabel').innerHTML = marker.getTitle();
  google.maps.event.addListener(marker, 'click', function () {
      getStair(this.id,this);
      getPreviewStair(this.id);
      document.getElementById('modalUser').value = this.creator;
      document.getElementById('modalStair').value = this.id;
  });
}

function appendToMarker(data, marker){
  document.getElementById('myModalLabel').innerHTML = data.stairname;
  
    document.getElementById('stairPhoto').src = marker.photo;
  
  
  document.getElementById('stairDesc').innerHTML = data.description;
  FB.api('/me', function(response) {
    document.getElementById('idtoken').value = response.id;    
    document.getElementById('idstair').value = $('me').value;
  });
  getComments(marker.id);
  $('#modal2').on('show.bs.modal', function () {
  $('.modal-content').css('height',515);
  //$('.modal-content').css('max-height','500px');
    });
  $('#modal2').modal('show');
  //('#modal-body').css('height',$('#modal-content').height()*0.50);
}


var appended = 0;
var papaDiv;
function getUserFromComment(comments){
  
  if(comments == null){
    $("#addCommentsHere").empty();
      return;
  }
  papaDiv = document.getElementById('addCommentsHere');
  papaDiv.style.cssText='text-align:center;overflow:auto;';
  if(appended == 0){
      document.getElementById('commentCont').appendChild(papaDiv);
      document.getElementById('comments').appendChild(document.getElementById('commentCont'));
      document.getElementById('tabcontent').appendChild(document.getElementById('comments'));
      document.getElementById('modalCont').appendChild(document.getElementById('tabcontent'));
        appended = 1;
        console.log('appeneded');
  }
  $("#addCommentsHere").empty();
  //document.removeChild();
 
  for(var i = 0; i < comments.length; i++){
    getUser(comments[i].idToken,comments[i], 'comment');
  }

}

function changeModalInput(tab){
    if(tab == 'comments'){
      document.getElementById('commentFormDiv').style.display = 'block';
      document.getElementById('uploadFormDiv').style.display = 'none';
    }else if(tab == 'photos'){
      document.getElementById('commentFormDiv').style.display = 'none';
      document.getElementById('uploadFormDiv').style.display = 'block';

    }else{
      document.getElementById('commentFormDiv').style.display = 'none';
      document.getElementById('uploadFormDiv').style.display = 'none';
    }
}

function createPhotos(photos){
  if(photos == null){
      document.getElementById('photos').innerHTML = "<p id='nopics'>There are no picture for this location</p>";
      return;
  }
   $('#photos').empty();

  var mamaDiv = document.getElementById('photos');

  for(var i = 0; i < photos.length; i++){
    var img = document.createElement('img');
    img.style.cssText = ' margin: 5px;'
    img.id = photos[i].id;
    img.src = photos[i].preview;
    mamaDiv.appendChild(img);
  }
}

function createComments(comment, user){
 
      var c = document.createElement('div');
      c.className = 'row';
      c.style.cssText = 'border: 1px solid;width:80%; position:relative; text-align:left; margin:0 auto;';
      var p2 = document.createElement('p');
      p2.style.cssText = "font-size:14px; float:left; width:49%;  margin-left:2px;";
      p2.id = 'commenter'+comment.commentId;
      p2.innerHTML = user.first_name + " "+ user.last_name;
      c.appendChild(p2);
      var p3 = document.createElement('p');
      p3.style.cssText = "font-size:14px; width:49%; float:right; text-align:right; margin-right:2px;";
      p3.id = 'time'+comment.commentId;
      p3.innerHTML = comment.commentDate.replace('T',' ').replace('+01:00','')+"<br>";
      c.appendChild(p3);
      var p1 = document.createElement('p');
      p1.style.cssText = "font-size:16px; width:100%; height:50%; margin:2px";
      p1.id = 'comtext'+comment.commentId;
      p1.innerHTML = comment.commentText;
      c.appendChild(p1);
      
      papaDiv.appendChild(c);

  

}
