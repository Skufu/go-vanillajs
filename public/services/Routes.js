import { HomePage } from "../components/HomePage";
import { MovieDetailsPage } from "../components/MovieDetailsPage";

export const routes = [
    {
        path: "/",
        component: HomePage
    },
    {
        path: /\/movies\/(\d+)/,
        component: MovieDetailsPage
    },
    {
        path: "movies",
        component: MoviesPage

    }

]