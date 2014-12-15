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
      creator:location.user,
      avg: location.average

  });   
  document.getElementById('myModalLabel').innerHTML = marker.getTitle();
  google.maps.event.addListener(marker, 'click', function () {
      getStair(this.id,this);
      getPreviewStair(this.id);
      document.getElementById('originalPic').style.display='none';
      document.getElementById('previews').style.display='block';
      document.getElementById('modalUser').value = this.creator;
      document.getElementById('modalStair').value = this.id;
  });
}

function showOriginal(data) {
    img = new Image();
    img.src = data.photo;
    var c = document.getElementById("originalPic");
    var ctx = c.getContext("2d");
    c.height = 270;
    c.width = 566;
    console.log('SIZE:' + c.width + " " + c.height);
    ctx.drawImage(img,0,0, c.width, c.height);
    c.style.display='block';
    document.getElementById('previews').style.display='none';
    document.getElementById('modalFooter').style.display='none';

}
function hideCanvas(){
    document.getElementById('previews').style.display='block';
    var c = document.getElementById("originalPic");
    document.getElementById('modalFooter').style.display='block';
    c.style.display = 'none';
}

function appendToMarker(data, marker){
  console.log(data);
  data.photo = marker.photo;
  document.getElementById('myModalLabel').innerHTML = data.stairname; 
  document.getElementById('rating').innerHTML = "Rating: "+ data.average;
  document.getElementById('stairPhoto').src = data.photo;
  
  
  document.getElementById('stairDesc').innerHTML = data.description;
  FB.api('/me', function(response) {
    document.getElementById('idstair').value = data.id;    
    document.getElementById('idtoken').value = response.id;
  });
  getComments(data.id);
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
      if(document.getElementById('previews').style.display=='none'){
            document.getElementById('modalFooter').style.display='block';
            document.getElementById('originalPic').style.display='none';
      }
    }else if(tab == 'photos'){
      document.getElementById('commentFormDiv').style.display = 'none';
      document.getElementById('uploadFormDiv').style.display = 'block';
        if(document.getElementById('previews').style.display=='none'){
            document.getElementById('modalFooter').style.display='none';
            document.getElementById('originalPic').style.display='block';
      }
    }else{
      document.getElementById('commentFormDiv').style.display = 'none';
      document.getElementById('uploadFormDiv').style.display = 'none';
      if(document.getElementById('previews').style.display=='none'){
            document.getElementById('modalFooter').style.display='block';
              document.getElementById('originalPic').style.display='none';
      }
    }
}

function createPhotos(photos){
  if(photos == null){
      document.getElementById('previews').innerHTML = "<p id='nopics'>There are no picture for this location</p>";
      return;
  }
   $('#previews').empty();

  mamaDiv = document.getElementById('previews');
  //document.getElementById('photos').appendChild(mamaDiv);
  for(var i = 0; i < photos.length; i++){
    var img = document.createElement('img');
    img.style.cssText = ' margin: 5px;'
    img.id = photos[i].photoId;
    img.src = photos[i].preview;
    img.onclick = function(){
      getOriginal(this);
    };
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
