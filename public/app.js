import { API } from "./services/API.js";
import HomePage from "./components/HomePage.js";
import AnimatedLoading from "./components/AnimatedLoading.js";
import MovieDetailsPage from "./components/MovieDetailsPage.js";
import YoutubeEmbed from "./components/YoutubeEmbed.js";



window.addEventListener("DOMContentLoaded", event => {
    // Check if we're on a movie details page (has movie ID in URL)
    const urlParams = new URLSearchParams(window.location.search);
    const movieId = urlParams.get('id');

    if (movieId) {
        // Show movie details page
        document.querySelector("main").appendChild(new MovieDetailsPage());
    } else {
        // Show home page
        document.querySelector("main").appendChild(new HomePage());
    }
});

window.app = {
    search: (event) => {
        event.preventDefault();
        const q = document.querySelector("input[type=search]").value;
        //TODO: Implement the search
        
    },
    api: API
}