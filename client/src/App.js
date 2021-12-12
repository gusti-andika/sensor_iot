import React, {useState, useEffect} from "react";
import axios from 'axios'

export default () => {
    const [latest, setLatest] = useState(-1)
    const [color, setColor] = useState('')
    const fetchLatest = async () => {
        const result = await axios.get('http://localhost:8000/sensor/latest')
        setLatest(result.data.latest)

        if (result.data.latest <= 30)
            setColor('green')
        else
            setColor('red')
        
    }

    useEffect(() => {
        const interval = setInterval(() => {
            fetchLatest()
        }, 2000)
        return () => clearInterval(interval)
    })

    return (<div className="container">
        <label>Temperature: </label>
            <span style={{backgroundColor: `${color}`, margin:"5px", padding: '0px 10px', width: "10px", fontWeight:"bolder"}}>
                {latest}
            </span>
        </div>)
}