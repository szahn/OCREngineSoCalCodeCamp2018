import React, {Component} from 'react';
import Document from './document';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';
import { withStyles } from '@material-ui/core/styles';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import UploadForm from './uploadForm';

const styles = theme => ({
    card: {
        margin: theme.spacing.unit * 2,
    },
    button: {
        margin: theme.spacing.unit,
    },
  });

class App extends Component {

    constructor(props){
        super(props);

        const {classes} = props;
        this.classes = classes;

        this.state = {
            coords: null,
            imageFilename: null,
            dialogMessage: null,
            showDialg: false
        }
        
        this.onDialogClose = this.onDialogClose.bind(this);
        this.onError = this.onError.bind(this);
        this.onOCR = this.onOCR.bind(this);
    }

    onDialogClose() {
        this.setState({ showDialg: false });
    };

    onError(message){
        debugger;
        this.setState({
            dialogMessage: message,
            showDialg: true
        });
    }

    onOCR(data){
        this.setState({
            coords: data.coords,
            imageFilename: data.imageFilename
        });
    }

    render(){
        return <Grid container direction='column' style={{flexGrow: 1}}>
            <Grid item>
                <UploadForm onOCR={this.onOCR} onError={this.onError} apiUrl="http://localhost:8282"/>
            </Grid>
            <Grid item>
                {this.state.imageFilename && <Document coords={this.state.coords} imageFilename={this.state.imageFilename}/>}
            </Grid>
            <Dialog
                open={this.state.showDialg}
                keepMounted
                onClose={this.onDialogClose}
                aria-labelledby="alert-dialog-slide-title"
                aria-describedby="alert-dialog-slide-description">
                    <DialogTitle id="alert-dialog-slide-title">
                        {"Error"}
                    </DialogTitle>
                    <DialogContent>
                        <DialogContentText id="alert-dialog-slide-description">
                        {this.state.dialogMessage}
                        </DialogContentText>
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={this.onDialogClose} color="primary">
                        Close
                        </Button>
                    </DialogActions>
            </Dialog>            
        </Grid>
    }

}

export default withStyles(styles)(App);
