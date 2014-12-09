var currentMark
var loc;
function createMarker(location) {
  var marker = new google.maps.Marker({
      id: location.id,
      position: location.position,
      map: map,
      title: location.stairname,
      description: location.description,
      photo: location.photo
  });   
  loc = location;
  var contentString;
  var infowindow = new google.maps.InfoWindow({
    content: contentString
  });
  var com = comments;
  google.maps.event.addListener(marker, 'click', function () {
          document.getElementById('myModalLabel').innerHTML = this.getTitle();
          document.getElementById('stairPhoto').src = this.photo;
          document.getElementById('stairDesc').innerHTML = this.description;
          document.getElementById('idstair').value = this.id;
          FB.api('/me', function(response) {
            document.getElementById('idtoken').value = response.id;    
          });
          getComments(this.id);
          $('#modal2').on('show.bs.modal', function () {
          $('.modal-content').css('height',$( window ).height()*0.9);
          //$('.modal-content').css('max-height','500px');
            });
          $('#modal2').modal('show');
  });
}
var papaDiv;
function getUserFromComment(comments){
  
  if(comments == null){
    $("#addCommentsHere").empty();
      return;
  }
  $("#addCommentsHere").empty();
  //document.removeChild();
  papaDiv = document.getElementById('addCommentsHere');
  papaDiv.style.cssText='text-align:center;overflow:auto;';
  document.getElementById('commentDiv').appendChild(papaDiv);
  for(var i = 0; i < comments.length; i++){
    getUser(comments[i].idToken,comments[i], 'comment');
  }

}

function createComments(comment, user){
 
      var c = document.createElement('div');
      c.className = 'row';
      c.style.cssText = 'height:20%; width:80%; position:relative;';
      var p2 = document.createElement('p');
      p2.style.cssText = "font-size:14px;";
      p2.id = 'commenter'+comment.commentId;
      p2.innerHTML = user.first_name + user.last_name + " - " + comment.commentDate.replace('T',' ').replace('+01:00','');
      c.appendChild(p2);
      var p1 = document.createElement('p');
      p1.style.cssText = "font-size:16px;";
      p1.id = 'comtext'+comment.commentId;
      p1.innerHTML = comment.commentText;
      c.appendChild(p1);

      
      papaDiv.appendChild(c);

  

}
