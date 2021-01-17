import React from "react";
import {List} from "react-content-loader";

export default class ListPlaceholder extends React.Component {
    render() {
        const number = this.props.number || 5;
        let list = [];

        for (let i = 0; i < number; i++) {
            list.push(<List key={i} className="mt-2 mb-2"/>)
        }

        return (
            <>{list}</>
        );
    }
}