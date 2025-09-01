export const API = {
    baseURL: "/api/",
    getTopMovies: async () => {
        return await API.fetch("movies/top/");
    },
    getRandomMovies: async () => {
        return await API.fetch("movies/random/");
    },
    getMovieById: async (id) => {
        return await API.fetch(`/movies/${id}`);
    },
    searchMovies: async (q, order, genre) => {
        return await API.fetch("movies/search/", {q, order, genre});
    },
    getGenres: async () => {
        return await API.fetch("genres/");
    },
    fetch: async (serviceName, args) => {
        try{
            
            // Example: "movies/search/" + "?q=batman&order=score&genre=1"
            const queryString = args ? new URLSearchParams(args).toString() : "";

            
            // Example: "/api/movies/search/?q=batman&order=score&genre=1"
            const response = await fetch(API.baseURL + serviceName + '?' + queryString);

            //  CONVERT TO JSON
            const result = await response.json();

            //  RETURN THE DATA (this is what gets returned!)
            return result;
        } catch (e) {
            console.error(e);
            app.showError("Error fetching data");
        }
    }
}