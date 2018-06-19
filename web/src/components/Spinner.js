import React from 'react';
import loading from "../images/loading.svg"

const style = {
    display: "block",
    marginTop: "15%",
    marginRight: "auto",
    marginLeft: "auto",
    maxWidth: "200px",
    maxHeight: "200px",
    width: "40%",
};

export default ({}) => {
    return <img style={style} src={loading} alt="Loading..."/> ;
};