import React, {Component} from 'react';
import {Redirect} from "react-router";
import Interactive from "react-interactive";

const movieCardOnHoverStyle = {
    boxShadow: "0 20px 25px 0 rgba(0,0,0,0.19), 0 20px 25px 0 rgba(0,0,0,0.19)",
};

const movieCardStyle = {
    boxShadow: "0 16px 18px 0 rgba(0,0,0,0.12), 0 19px 25px 0 rgba(0,0,0,0.09)",
    background: "#fff",
    transitionDuration: "0.3s",
    padding:"1px 20px",
    margin: "20px 0px",
    maxWidth: "950px",
    marginLeft: "auto",
    marginRight: "auto",
};


class MovieList extends Component {

    constructor(props){
        super(props);
        this.state = {
            redirect: null,
        }
    }

    render() {
        if (this.state.redirect) {
            return <Redirect push to={"/" + this.state.redirect}/>;
        }
        let rows = [];
        for (let i in this.props.years) {
           const year = this.props.years[i];
           rows.push(
               <Interactive
                   as="div"
                   key={i}
                   hover={movieCardOnHoverStyle}
                   style={movieCardStyle}
                   onClick={() => this.setState({redirect: year.year})}>
                   <h2 >{year.name}<br/><i>{"eller " + year.eller}</i> ({year.year})</h2>
                   <table key={"table" + i} style={{width: "100%"}}>
                       <tbody>
                       <tr>
                           <td width="200px">
                               <img width="100%" src={year.poster_uri} alt="poster"/>
                           </td>
                           <td style={{verticalAlign: "top", padding:"10px 20px"}}>
                               {year.description}
                           </td>
                       </tr>
                       </tbody>
                   </table>
               </Interactive>
           );
        }
        return rows
    }
}

export default MovieList

export const MovieCardStyle = movieCardStyle;