import reviews
import genres

PICTURES_FUNCTIONS_MAP = {
    "genres_piechart": genres.genres_pie,
    "total_games": genres.genres_addition,
    "relative_genres": genres.genres_addition,
    "total2_games": genres.genres_detailed,
    "relative2_genres": genres.genres_detailed,
    "reviews": reviews.reviews_distrib,
    "timeline_median_reviews": reviews.reviews_timeline,
    "timeline_avg_reviews": reviews.reviews_timeline
}

def generate_mapping():
    for gnr in genres.GENRES_NAMES:
        PICTURES_FUNCTIONS_MAP["paper_" + gnr] = genres.genres_paper
        PICTURES_FUNCTIONS_MAP["relative_" + gnr] = genres.genres_addition
        PICTURES_FUNCTIONS_MAP["detailed2_" + gnr] = genres.genres_detailed
        PICTURES_FUNCTIONS_MAP["reltive2_" + gnr] = genres.genres_detailed
        PICTURES_FUNCTIONS_MAP["reviews_" + gnr] = reviews.reviews_genres

