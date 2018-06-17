import React, {Component} from 'react';
import jwtDecode from 'jwt-decode'
import Api from "../api";

const headerStyle = {
    margin: "20px",
};

const MovieHeader = (props) => {
    return null;
};

const Movies = (props) => {
    const views = [];
    for (let i in props.movies) {
        const movie = props.movies[i];
        views.push(
            <div key={i}>
                <video controls mediaGroup="video" src={movie.uri} style={{width:"200px"}}/>
                <p>
                    {movie.description}
                </p>
            </div>
        )
    }
    return views;
};

class MovieView extends Component{

    constructor(props) {
        super(props);
        this.state = {
            user: {},
            tokenData: {},
            movies: [],
        }
    }

    componentDidMount() {
        try {
            const tokenData = jwtDecode(this.props.token);
            this.setState({tokenDate: tokenData});
            Api.GetMovies(this.props.year, this.props.token)
                .then(movies => this.setState({movies: movies}));
        } catch(error) {
            console.log(error)
        }
    }

    render() {
        return (
            <div style={headerStyle}>
                <MovieHeader token={this.props.token} year={this.props.year}/>
                <p>{this.props.description}</p>
                <Movies movies={this.state.movies}/>
            </div>
        )
    }
}

export default MovieView