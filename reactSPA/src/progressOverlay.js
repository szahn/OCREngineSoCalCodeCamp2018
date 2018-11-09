import React from 'react';

import CircularProgress from '@material-ui/core/CircularProgress';

const styles = {
    progressBackground:{
        background: `rgba(255,255,255,.5)`,
        position:'absolute',
        left:0,
        top:0,
        right:0,
        bottom:0,
        zIndex: 999,
        textAlign: 'center',
        margin: '0 auto',
        cursor: 'wait'
    },
    progress:{
        margin: '20px auto'
    }
}

export default ()=> <div style={styles.progressBackground}>
    <CircularProgress style={styles.progress}/>
</div>