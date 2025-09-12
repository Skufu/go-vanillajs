const Router = {
   init: () => {
    window.addEventListener("popstate", () => {
        Router.go(location.pathname, false);
    });

    //go to the intial route
    Router.go(location.pathname + location.search);
   },

   go: (route, addToHistory=true) => {
    if (addToHistory) {
        history.pushState(null, "", route);
    }


   }
}