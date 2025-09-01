import { API } from "./services/API.js";
import { HomePage } from "./components/HomePage.js";
import './components/AnimatedLoading.js';

window.addEventListener("DOMContentLoaded", event => {
    document.querySelector("main").appendChild(new HomePage());
});

window.app = {
    search: (event) => {
        event.preventDefault();
        const q = document.querySelector("input[type=search]").value;
        //TODO: Implement the search
        
    },
    api: API
}