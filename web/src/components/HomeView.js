import React, {Component} from 'react';
import Api from "../api";
import MovieList from "./MovieList";
import Spinner from "./Spinner";

const headerStyle = {
    margin: "20px",
};

class HomeView extends Component{

    constructor(props) {
        super(props);
        this.state = {
            years: [],
        }
    }

    componentDidMount() {
        Api.GetAllYears(this.props.token)
            .then(years => years.sort((y1, y2) => y2.year - y1.year))
            .then(years => this.setState({years: years}));
    }

    render() {
        return (
            <div style={headerStyle}>
            {this.state.years.length > 0
                ? <MovieList years={this.state.years}/>
                : <Spinner/> }
            </div>
        )
    }
}

export default HomeView