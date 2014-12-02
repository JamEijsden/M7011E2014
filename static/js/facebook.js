// This is called with the results from from FB.getLoginStatus().
  function statusChangeCallback(response) {
    console.log('statusChangeCallback');
    console.log(response);
    //var div_FBLogin = document.getElementById('fb-login');
    //var div_FBLike = document.getElementById('fb-liike');
    // The response object is returned with a status field that lets the
    // app know the current login status of the person.
    // Full docs on the response object can be found in the documentation
    // for FB.getLoginStatus().
    if (response.status === 'connected') {
      // Logged into your app and Facebook.
    //  div_FBLogin.style.display = 'inline';
     // div_FBLike.style.display = 'inline';
      //replace_login(true);
      if(document.location.href == "http://trollegeuna.se:9999/") {
		document.location.href = "http://trollegeuna.se:9999/test/";
	}
      testAPI();

    } else if (response.status === 'not_authorized') {
      // The person is logged into Facebook, but not your app.
      //div_FBLogin.style.display = 'inline';
     // div_FBLike.style.display = 'none';
      //replace_login(false);

    } else {
      // The person is not logged into Facebook, so we're not sure if
      // they are logged into this app or not.
      //replace_login(false);
      //div_FBLogin.style.display = 'inline';
     // div_FBLike.style.display = 'none';

    }
  }

  // This function is called when someone finishes with the Login
  // Button.  See the onlogin handler attached to it in the sample
  // code below.

  function checkLoginState() {
    FB.getLoginStatus(function(response) {
      statusChangeCallback(response);
    });
  }

  function logout(){
      FB.logout() 
      window.location.reload();
  }
  function login() {
      FB.login(function(response) {
          if (response.authResponse) {
              // connected

	    checkLoginState();
	      window.location.href = "http://trollegeuna.se:9999/test/";
          } else {
              // cancelled
              replace_login(false);
		alert('User cancelled login or did not fully authorize.');
          }
      });
  }

  /*function replace_login(logged_in){
    var elem = document.getElementById("fb-login-btn");
    if ((elem.innerHTML=="Login with Facebook") && logged_in){
	elem.innerHTML = "Logout";
	elem.onclick = function() { logout(); };
    } else {
 	elem.innerHTML = "Login with Facebook";
	elem.onlick = function() { login(); };		
    }
  }
*/
  window.fbAsyncInit = function() {
  FB.init({
    appId      : '562407890559656',
    cookie     : true,  // enable cookies to allow the server to access 
                        // the session
    xfbml      : true,  // parse social plugins on this page
    version    : 'v2.1' // use version 2.1
  });

  // Now that we've initialized the JavaScript SDK, we call 
  // FB.getLoginStatus().  This function gets the state of the
  // person visiting this page and can return one of three states to
  // the callback you provide.  They can be:
  //
  // 1. Logged into your app ('connected')
  // 2. Logged into Facebook, but not your app ('not_authorized')
  // 3. Not logged into Facebook and can't tell if they are logged into
  //    your app or not.
  //
  // These three cases are handled in the callback function.

  FB.getLoginStatus(function(response) {
    statusChangeCallback(response);
  });

  };

  // Load the SDK asynchronously
  (function(d, s, id) {
    var js, fjs = d.getElementsByTagName(s)[0];
    if (d.getElementById(id)) return;
    js = d.createElement(s); js.id = id;
    js.src = "//connect.facebook.net/en_US/sdk.js";
    fjs.parentNode.insertBefore(js, fjs);
  }(document, 'script', 'facebook-jssdk'));

  // Here we run a very simple test of the Graph API after login is
  // successful.  See statusChangeCallback() for when this call is made.
  function testAPI() {
   // document.location = "localhost:9999/about/";
    console.log('Welcome!  Fetching your information.... ');
    FB.api('/me', function(response) {
      console.log('Successful login for: ' + response.name);
     //   document.getElementById('photo').src="https://graph.facebook.com/"+response.id+"/picture"; 
    });
  }
