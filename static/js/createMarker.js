var currentMark
var loc;
function createMarker(location, comments) {
  var marker = new google.maps.Marker({
      id: location.id,
      position: location.position,
      map: map,
      title: location.stairname,
      description: location.description,
      photo: location.photo
  });   
  console.log("Created: " + marker.position);
  loc = location;
  var contentString;
  console.log(loc);
  var infowindow = new google.maps.InfoWindow({
    content: contentString
  });
  console.log(comments);
  var com = comments;
  google.maps.event.addListener(marker, 'click', function () {
          document.getElementById('myModalLabel').innerHTML = this.getTitle();
          document.getElementById('stairPhoto').src = this.photo;
          document.getElementById('stairDesc').innerHTML = this.description;
          document.getElementById('idstair').value = this.id;
          FB.api('/me', function(response) {
            console.log('user commenting:  ' + response.id);
            document.getElementById('idtoken').value = response.id;
            getUser(response.id, com,'comment');
          });
          $('#modal2').on('show.bs.modal', function () {
          $('.modal-content').css('height',$( window ).height()*0.9);
          $('.modal-content').css('max-height','500px');
            });
          $('#modal2').modal('show');
  });
}

function createComments(comments, user){
 $("#addCommentsHere").remove();
  //document.removeChild();
  var papaDiv = document.createElement('div');
  papaDiv.id = 'addCommentsHere';
  papaDiv.style.cssText='text-align:center;';
  document.getElementById('commentDiv').appendChild(papaDiv);
   for (var i=0; i < comments.length; i++) {
      console.log(comments[i]);
      var c = document.createElement('div');
      c.className = 'row';
      c.style.cssText = 'height:20%; width:80%; position:relative;';
      var p2 = document.createElement('p');
      p2.style.cssText = "font-size:14px;";
      p2.id = 'commenter'+i;
      p2.innerHTML = user.first_name + user.last_name + " - " + comments[i].commentDate;
      c.appendChild(p2);
      var p1 = document.createElement('p');
      p1.style.cssText = "font-size:16px;";
      p1.id = 'comtext'+i;
      p1.innerHTML = comments[i].commentText;
      c.appendChild(p1);

      
      papaDiv.appendChild(c);

    }

}
