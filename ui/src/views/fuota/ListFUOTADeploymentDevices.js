import React, { Component } from "react";

import FUOTADeploymentStore from "../../stores/FUOTADeploymentStore";


class ListFUOTADeploymentDevices extends Component {
  constructor() {
    super();
    this.state = {};

    this.getPage = this.getPage.bind(this);
    this.getRow = this.getRow.bind(this);
  }

  getPage(limit, offset, callbackFunc) {
    FUOTADeploymentStore.list
  }

  componentDidMount() {
    FUOTADeploymentStore.get(this.props.match.params.fuotaDeploymentID, resp => {
      this.setState({
        fuotaDeployment: resp,
      });
    });
  }
}
