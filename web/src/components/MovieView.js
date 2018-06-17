import React, {Component} from 'react';
import jwtDecode from 'jwt-decode'
import Api from "../api";

const movieHeaderStyle = {
    width: "100%",
    height: "350px",
    maxHeight: "50%",
    objectFit: "cover",
    objectPosition: "50% 10%",
};

const gradientStyle = {
    width: "100%",
    height: "100%",
    position: "absolute",
    bottom: "0px",
    left: "0px",
    background: "linear-gradient(rgba(0,0,0,0), black)",
};

const titleStyle = {
    after:"",
    position: "absolute",
    bottom: "8px",
    left: "16px",
    fontSize: "2em",
    color: "#dbcd30",
};

const MovieHeader = (props) => {
    console.log(props);
    return (
        <div style={{position: "relative"}}>
            <img src={props.yearData.poster_uri} style={movieHeaderStyle} alt={props.yearData.name}/>
            <div style={gradientStyle}/>
            <div style={titleStyle}>
                {props.yearData.name} <br/>
                <i>eller {props.yearData.eller}</i>
            </div>
        </div>
);
};

const movieStyle = {
    width: "200px",
    boxShadow: "0 16px 18px 0 rgba(0,0,0,0.12), 0 19px 25px 0 rgba(0,0,0,0.09)",
};

const Movies = (props) => {
    const views = [];
    for (let i in props.movies) {
        const movie = props.movies[i];
        views.push(
            <div key={i} style={{display: "flex", flexDirection: "row"}}>
                <div >
                    <h3 style={{margin: "10px 0px"}}>{movie.name}</h3>
                    <video style={movieStyle} controls mediaGroup="video" src={movie.uri}/>
                    <p>
                        asdasdasd{movie.description}
                    </p>
                </div>
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
            yearData: {},
        }
    }

    componentDidMount() {
        try {
            const tokenData = jwtDecode(this.props.token);
            this.setState({tokenDate: tokenData});
            Api.GetAllYears(this.props.token)
                .then(allYears => {
                    return allYears.filter(year =>
                        year.year === "" + this.props.year)[0]
                })
                .then(year => {
                    this.setState({yearData: year})
                });
            Api.GetMovies(this.props.year, this.props.token)
                .then(movies => this.setState({movies: movies}));
        } catch(error) {
            console.log(error)
        }
    }

    render() {
        return (
            <div>
                <MovieHeader yearData={this.state.yearData}/>
                <div style={{margin: "0px 20px"}}>
                    <p>{this.state.yearData.description}</p>
                    <Movies movies={this.state.movies}/>
                </div>
            </div>
        )
    }
}

export default MovieView