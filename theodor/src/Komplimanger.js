import React, {Component} from 'react';
import Button from "./Button";

const komplimanger = [
    "Teatralisk",
    "Härlig",
    "Energisk",
    "Ordentlig",
    "Driven",
    "Omhändertagande",
    "Rolig",

    "Mysig",
    "Omtänksam",
    "Ledarpotential",
    "Ambitiös",
    "Noggrann",
    "Dansar bra",
    "Entusiastisk",
    "Rustik",
    "THEODOR MOLANDER"
];

class Komplimang extends Component {

    constructor(props) {
        super(props);
        this.state = {
            komp: 0,
            letters: "",
        };
        this.nextKomplimang = this.nextKomplimang.bind(this);
    }

    style = {
        padding: "20px",
        fontSize: "1.5em",
    };

    headerStyle = {
        fontWeight: "normal",
        fontFamily: "'Playfair Display', serif",
    };

    KomplimangStyle = {
        fontWeight: "bold",
        fontFamily: "'Playfair Display', serif",
    };

    KomplimangEndStyle = {
        fontWeight: "bold",
        color: "#ffffff",
        fontFamily: "'Playfair Display', serif",
    };

    getKomplimangStyle() {
        if (this.state.komp < 15) {
            return this.KomplimangStyle;
        } else {
            return this.KomplimangEndStyle;
        }
    }

    theodorStyle = {
        fontWeight: "light",
        fontFamily: "'Playfair Display', serif",
    };

    theodorEndStyle = {
        fontWeight: "light",
        marginTop: "-115px",
        fontFamily: "'Playfair Display', serif",
        transitionDuration: "0.6s",
    };

    getTheodorStyle() {
        if (this.state.komp < 15) {
            return this.theodorStyle;
        } else {
            return this.theodorEndStyle;
        }
    }

    getComplimang(i) {
        return komplimanger[i % komplimanger.length]
    }

    nextKomplimang() {
        const i = this.state.komp;
        if (i < 15) {
            let lastLetter = this.getComplimang(i)[0];
            if (lastLetter === "M") {
                lastLetter = "\nM";
            }
            this.setState({
                letters: this.state.letters + lastLetter,
                komp: i + 1,
            })
        }
    }

    render() {
        return (
            <div style={this.style}>
                <p style={this.headerStyle}> Du är: </p>
                <p style={this.getKomplimangStyle()}>{this.getComplimang(this.state.komp)}</p>
                <Button text="Klick på mig!" onClick={this.nextKomplimang}/>
                <p style={this.getTheodorStyle()}>{this.state.letters}</p>
            </div>
        );
    }
}

const Komplimanger = (props) => {
    return (<Komplimang/>
    );
};

export default Komplimanger