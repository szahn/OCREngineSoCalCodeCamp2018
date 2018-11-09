import React, {Component} from 'react';
import { Stage, Layer, Rect, Image, Text } from 'react-konva';
import Konva from 'konva';
import axios from 'axios';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import {throttle} from 'lodash';
import { withStyles } from '@material-ui/core/styles';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';


const styles = theme => ({
    card: {
        margin: theme.spacing.unit * 2,
        cursor:'crosshair'
    },
   document: {
        width: '400px',
        height: '400px',
    }
});

class Document extends Component {

    constructor(props){
        super(props);

        const {classes} = props;
        this.classes = classes;

        this.state = {
            zoom: .25,
            offsetX: 0,
            offsetY: 0,
            image: null,
            documentTitle: ""
        };

        this.isMouseDown = false;

        this.docRef = React.createRef();

        this.onCoordClick = this.onCoordClick.bind(this);
        this.onWheel = throttle(this.onWheel, 30).bind(this);
        this.onResize = this.onResize.bind(this);
        this.onMouseDown = this.onMouseDown.bind(this);
        this.onMouseMove = throttle(this.onMouseMove, 30).bind(this);
        this.onMouseUp = this.onMouseUp.bind(this);
        this.onMouseLeave = this.onMouseLeave.bind(this);

        window.addEventListener('resize', this.onResize);

        this.loadImage(props.imageFilename);
    }

    componentWillUnmount(){
        window.removeEventListener('resize', this.onResize);
    }

    onResize(){
        const {current} = this.docRef;
        this.setState({
            docWidth: .99 * (current.clientWidth),
            docHeight: 0.9*(window.innerHeight - current.offsetTop)
        });
    }

    componentDidUpdate(prevProps){
        if (this.props.imageFilename !== prevProps.imageFilename){
            this.loadImage(this.props.imageFilename);
        }
    }

    loadImage(imageName){
        const image = new window.Image();
        const filename = `./img/${imageName}`;
        console.log(`Loading image ${filename}...`)
        image.src = filename;
        image.onload = () => {
            this.setState({
                image: image
            });
        };
    }

    componentDidMount(){
        const {current} = this.docRef;
        this.setState({
            docWidth: .99 * (current.clientWidth),
            docHeight: 0.9*(window.innerHeight - current.offsetTop)
        });
    }

    onCoordClick(coord, e){
        this.setState({
            documentTitle: `Clicked: ${coord.text}`
        });
    }

    onWheel(e){
        const deltaY = e.evt.deltaY;
        const dY = deltaY / 1000;
        this.setState((s)=> {
            const newZoom = s.zoom + dY;
            const zoomDiff = newZoom - s.zoom;
            s.zoom = newZoom;
        return s;
        });
    }

    onMouseDown(e){
        this.isMouseDown = true;
        this.mouseAnchor = [e.evt.clientX, e.evt.clientY];
        this.mouseOff = [e.evt.clientX, e.evt.clientY];
    }

    onMouseMove(e){
        if (!this.isMouseDown) return;

        let diffX = (this.mouseAnchor[0] - e.evt.clientX) / this.state.zoom;
        let diffY = (this.mouseAnchor[1] - e.evt.clientY) / this.state.zoom;

        this.mouseAnchor = [e.evt.clientX, e.evt.clientY];

        requestAnimationFrame(()=>{
            this.setState((s)=>{
                s.offsetX += diffX;
                s.offsetY += diffY;
                return s;
            });
        });
    }

    onMouseUp(e){
        this.isMouseDown = false;
    }

    onMouseLeave(e){
        this.isMouseDown = false;
    }
      
    render(){
        const {image, zoom, offsetX, offsetY} = this.state;
        const {coords} = this.props;
        const margin = 10;
        return <Card className={this.classes.card}>
            <Toolbar>
                <Typography variant="h6" color="inherit">
                    {this.state.documentTitle}
                </Typography>
            </Toolbar>
            <div ref={this.docRef}>
                <CardContent>
                    <Stage width={this.state.docWidth - margin} height={this.state.docHeight} scaleX={zoom} scaleY={zoom} offsetX={offsetX} offsetY={offsetY} onWheel={this.onWheel} onMouseDown={this.onMouseDown} onMouseMove={this.onMouseMove} onMouseUp={this.onMouseUp} onMouseLeave={this.onMouseLeave}>
                        <Layer>
                            {image && <Image image={image}/>}
                            {coords && coords.map(coord => 
                                <Rect onClick={this.onCoordClick.bind(this, coord)} width={coord.x2 - coord.x1} height={coord.y2 - coord.y1} stroke={'red'} strokeWidth={1} x={coord.x1} y={coord.y1} />
                            )}
                        </Layer>
                    </Stage>
                </CardContent>
            </div>
        </Card>
    }

}

export default withStyles(styles)(Document);
