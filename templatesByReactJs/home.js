import * as React from "react";
import * as $ from "jquery";
/**
 * Created by Ryan on 16/11/21.
 */

var PoemBox = React.createClass({
    loadDataFromServer: function () {
        $.ajax(
            {
                url: this.props.url,
                dataType: 'json',
                cache: false,
                success: function (data) {
                    this.setState(data.data)
                }.bind(this),

                error: function (xhr, status, err) {
                    console.error(this.props.url, status, err.toString());
                }.bind(this)

            }
        );

    },

    getInitialState: function () {
        return {data: []}
    },
    componentDidMount: function () {
        this.loadDataFromServer()
        setInterval(this.loadDataFromServer, this.props.pollInterval)
    },

    render: function () {
        return (
            <div className="poemBox">
            </div >
        );
    }
});

ReactDOM.render(
    <PoemBox url="/" pollInterval="{2000}"/>,
    document.getElementById("content")
)