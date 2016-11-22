/**
 * Created by Ryan on 16/11/22.
 */
var ArticleBox = React.createClass({
        loadArticleFromServer: function () {
            $.ajax({
                url: this.props.url,
                dataType: 'json',
                contentType: 'application/json',
                cache: false,
                success: function (data) {
                    console.log(data)
                    this.setState({data: data})
                }.bind(this),
                error: function (xhr, status, err) {
                    console.error(this.props.url, status, err.toString());
                }.bind(this)
            });
        },

        getInitialState: function () {
            return {data: []}
        },
        componentDidMount: function () {
            this.loadArticleFromServer();
            // setInterval(this.loadArticleFromServer, this.props.pollInterval);
        },

        render: function () {
            return (
                <div className="articleBox">
                    <h1>Articles</h1>
                    <ArticleList data={this.state.data}/>
                </div>

            )
        }


    }
)


var ArticleList = React.createClass({
        render: function () {
            var articleNodes = this.props.data.map(
                function (article) {
                    return (
                        <Article title={article.title} content={article.content} id={article.id}>
                            {/*{article.content}*/}
                        </Article>
                    )

                }
            );

            return (
                <div className="articleList">
                    {articleNodes}
                </div>

            )
        }

    }
)

var Article = React.createClass({
        render: function () {
            var url = `/article/${this.props.id}`;
            return (
                <div className="article">

                    <a className="title" href={url}>{this.props.title}</a>
                    <p className="content">{this.props.content}</p>
                </div>
            );
        }
    }
)

ReactDOM.render(
    <ArticleBox url="/api/articles"/>,
    document.getElementById("content")
);