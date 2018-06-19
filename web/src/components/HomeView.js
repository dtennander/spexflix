import React, {Component} from 'react';
import Api from "../api";
import MovieList, {MovieCardStyle} from "./MovieList";
import Spinner from "./Spinner";
import InputField from "./InputField";

const headerStyle = {
    margin: "20px",
};

const SearchStyle = {
    ...MovieCardStyle,
    height: "45px",
    width: "100%",
    padding: "0px 20px",
    display: "block",
    fontSize: "1.2em",
    border: "1px solid #eee",
    borderRadius: "8px",
    boxSizing: "border-box",
};

class HomeView extends Component{

    constructor(props) {
        super(props);
        this.state = {
            years: [],
            filteredList: [],
        };
        this.searchField = React.createRef();
        this.onSearchChange = this.onSearchChange.bind(this);
        this.filterSearch = this.filterSearch.bind(this);
    }

    componentDidMount() {
        Api.GetAllYears(this.props.token)
            .then(years => years.sort((y1, y2) => y2.year - y1.year))
            .then(years => this.setState({years: years, filteredList: years}));
    }

    onSearchChange(key) {
        if (key === "escape") {
            this.searchField.current.clearField();
            this.setState({filteredList: this.state.years});
        }
        if (key === "enter") {
            this.filterSearch(this.searchField.current.getInput());
        }
    }

    filterSearch(searchTerm) {
        console.log(searchTerm);
        let filteredList = this.state.years
            .filter(year =>
                Object.values(year)
                    .some(value => value.toLowerCase().includes(searchTerm.toLowerCase())));
        console.log(filteredList);
        this.setState({filteredList:filteredList});
    }

    render() {
        return (
            <div style={headerStyle}>
            {this.state.years.length > 0
                ? <div>
                    <InputField
                        ref={this.searchField}
                        style={SearchStyle}
                        type="text"
                        placeholder="Sök här"
                        onEnter={() => this.onSearchChange("enter")}
                        onEscape={() => this.onSearchChange("escape")}/>
                    <MovieList years={this.state.filteredList}/>
                </div>
                : <Spinner/> }
            </div>
        )
    }
}

export default HomeView