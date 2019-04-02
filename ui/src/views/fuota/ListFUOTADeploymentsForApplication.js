import React, { Component } from "react";
import { Link } from "react-router-dom";

import { withStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";
import Button from '@material-ui/core/Button';

import moment from "moment";
import CloudUpload from "mdi-material-ui/CloudUpload";

import TableCellLink from "../../components/TableCellLink";
import DataTable from "../../components/DataTable";
import Admin from "../../components/Admin";
import FUOTADeploymentStore from "../../stores/FUOTADeploymentStore";
import theme from "../../theme";


const styles = {
  buttons: {
    textAlign: "right",
  },
  button: {
    marginLeft: 2 * theme.spacing.unit,
  },
  icon: {
    marginRight: theme.spacing.unit,
  },
};


class ListFUOTADeploymentsForApplication extends Component {
  constructor() {
    super();

    this.state = {};

    this.getPage = this.getPage.bind(this);
    this.getRow = this.getRow.bind(this);
  }

  getPage(limit, offset, callbackFunc) {
    FUOTADeploymentStore.list({
      applicationID: this.props.match.params.applicationID,
      limit: limit,
      offset: offset,
    }, callbackFunc);
  }

  getRow(obj) {
    const createdAt = moment(obj.createdAt).format('lll');
    const updatedAt = moment(obj.updatedAt).format('lll');

    return(
      <TableRow key={obj.id}>
        <TableCellLink to={`/organizations/${this.props.match.params.organizationID}/applications/${this.props.match.params.applicationID}/fuota-deployments/${obj.id}`}>{obj.name}</TableCellLink>
        <TableCell>{createdAt}</TableCell>
        <TableCell>{updatedAt}</TableCell>
        <TableCell>{obj.state}</TableCell>
      </TableRow>
    );
  }

  render() {
    return(
      <Grid container spacing={24}>
        <Admin organizationID={this.props.match.params.organizationID}>
          <Grid item xs={12} className={this.props.classes.buttons}>
            <Button variant="outlined" className={this.props.classes.button} component={Link} to={`/organizations/${this.props.match.params.organizationID}/applications/${this.props.match.params.applicationID}/fuota-deployments/create`}>
              <CloudUpload className={this.props.classes.icon} />
              Create FUOTA Deployment
            </Button>
          </Grid>
        </Admin>

        <Grid item xs={12}>
          <DataTable
            header={
              <TableRow>
                <TableCell>Name</TableCell>
                <TableCell>Created at</TableCell>
                <TableCell>Updated at</TableCell>
                <TableCell>State</TableCell>
              </TableRow>
            }
            getPage={this.getPage}
            getRow={this.getRow}
          />
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(ListFUOTADeploymentsForApplication);

