import { API } from "./services/API.js";
import HomePage from "./components/HomePage.js";
import AnimatedLoading from "./components/AnimatedLoading.js";
import MovieDetailsPage from "./components/MovieDetailsPage.js";
import YoutubeEmbed from "./components/YoutubeEmbed.js";
import { Router } from "./services/Router.js";



window.addEventListener("DOMContentLoaded", event => {
    app.Router.init();
});

window.app = {
    Router,
    search: (event) => {
        event.preventDefault();
        const q = document.querySelector("input[type=search]").value;
        //TODO: Implement the search
        
    },
    api: API
}