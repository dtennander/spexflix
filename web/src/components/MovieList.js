import React, {Component} from 'react';
import Api from "../api";


const playerStyle = {
    maxWidth: "800px",
};

class MovieList extends Component {

    constructor(props){
        super(props);
        this.state = {
            movies: []
        }
    }

    componentDidMount() {
        Api.GetAllMovies(this.props.token)
            .then(movies => {
                this.setState({movies: movies});
            })
    }

    render() {
        let rows = [];
        console.log(this.state.movies);
        for (let year in this.state.movies) {
            let movies = this.state.movies[year];
            if (movies == null) {
                continue
            }
            console.log(movies);
            rows.push(<h2 key={year}>{year}</h2>);
            for (let i in movies) {
                const movie = movies[i];
                rows.push(
                    <div key={year + ":" + i} style={playerStyle}>
                        <h3>{movie.Name}</h3>
                        <video controls mediaGroup="video" src={movie.Uri} style={{width:"100%"}}/>
                    </div>
                );
            }
        }
        return (
            <div>
                {rows}
            </div>
        )
    }
}

export default MovieList