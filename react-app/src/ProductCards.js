import React from 'react';

// Card component for a single product
class Card extends React.Component {
  render() {
    const priceColor = (this.props.promo)? "text-danger" : "text-dark";            // If the product has promotion, show price in red. Otherwise show in black.
    const sellPrice  = (this.props.promo)? this.props.promotion :this.props.price  // If the product has promotion, show the promotion price. Otherwise show original price.
    return (
      <div className="col-md-6 col-lg-4 d-flex align-items-stretch">
        <div className="card mb-3">
          <img className="card-img-top" src={this.props.img} alt={this.props.imgalt} />
          <div className="card-body">
            <h4 className="card-title">{this.props.productname}</h4>
            Price : <strong className={priceColor}>{sellPrice}</strong>
            <p className="card-text">{this.props.desc}</p>
            <a className="btn btn-success text-white" onClick={()=>{this.props.showBuyModal(this.props.ID,sellPrice)}}>Buy</a>
          </div>
        </div>
      </div>
    );
  }
}

// CardContainer component for a list of products
export default class CardContainer extends React.Component {
  constructor(props) {
    super(props);
      this.state = {
        cards: []
      };
  }

  componentDidMount() {
    fetch(this.props.location)
        .then(res => res.json())
        .then((result) => {
          this.setState({
            cards: result
          });
        });
  }

  render() {
    const cards = this.state.cards;
    let items = cards.map(
      card => <Card  key={card.ID} {...card} promo={this.props.promo} showBuyModal={this.props.showBuyModal} />
    );
    return (
      <div>
        <div className="mt-5 row">
          {items}
        </div>
      </div>
    );
  }
}
