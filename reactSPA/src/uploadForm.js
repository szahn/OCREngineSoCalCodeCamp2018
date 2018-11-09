import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import ProgressOverlay from './progressOverlay';
import Typography from '@material-ui/core/Typography';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';

const styles = theme => ({
    card: {
        margin: theme.spacing.unit * 2,
        position:'relative'
    },
    button: {
        margin: theme.spacing.unit,
    },
  });


class UploadForm extends React.Component{

    constructor(props){
        super(props);

        const {classes} = props;
        this.classes = classes;

        this.state = {
            isBusy: false
        };
        
        this.onFormSubmit = this.onFormSubmit.bind(this);
    }

    onFormSubmit(e){
        e.preventDefault();

        if (!this.uploadInput.files.length){
            this.props.onError("No file uploaded.");
            return;
        }

        this.setState({
            isBusy: true
        });

        const data = new FormData();
        const file = this.uploadInput.files[0];
        data.append('file', file);
        data.append('filename', file.name);

        try{

            let timedOut = false;
            const timeoutHandle = setTimeout(() => {
                this.setState({
                    isBusy: false
                });
                
                timedOut = true;
                this.props.onError("An error occured while OCRing document");
            }, 10 * 1000);

            fetch(`${this.props.apiUrl}/upload`, {
                method: 'POST',
                body: data            
            })
            .then((response => response.json()))
            .then(json => {
                if (timedOut) return;
                this.props.onOCR({
                    coords: json.blocks,
                    imageFilename: json.imageName
                });
            }).catch(err=>{
                this.props.onError("An error occured while OCRing document");
            }).then(()=>{
                clearTimeout(timeoutHandle);
                this.setState({
                    isBusy: false
                });
            });
        }
        catch (err){
            clearTimeout(timeoutHandle);
            this.props.onError(`An error occured while OCRing document. ${err.message}`);
            this.setState({
                isBusy: false
            });
        }
    }

    render(){
        return <Card className={this.classes.card}>
            <CardContent>
                <Typography color="textSecondary" variant="title" gutterBottom>
                    Upload a Document
                </Typography>
            </CardContent>
            <CardActions>
                <form onSubmit={this.onFormSubmit}>
                    <input accept="application/pdf" type="file" ref={(r)=> {this.uploadInput = r;}} style={{ display: 'none' }} id="raised-button-file"/>
                    <label htmlFor="raised-button-file" className={this.classes.button}>
                        <Button size="small" variant="contained" color="secondary" component="span">Upload PDF</Button>                                        
                    </label>
                    <Button variant="contained" color="primary"  type="submit" size="small" className={this.classes.button}>
                        OCR
                    </Button>
                </form>
            </CardActions>
            {this.state.isBusy && <ProgressOverlay/>}
        </Card>;
    }
}

export default withStyles(styles)(UploadForm);
