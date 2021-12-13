import React, {useState, useEffect} from "react";
import axios from 'axios'

export default (prop) => {
    const [temp, setLatest] = useState(-1)
    const [color, setColor] = useState('')
    const fetchLatest = async () => {
        const result = await axios.get(`http://localhost:8000/sensor/${prop.kind}`)
        setLatest(result.data.value)

        if (prop.useColor) {
            if (result.data.value <= 30)
                setColor('green')
            else
                setColor('red')
        }
        
    }

    useEffect(() => {
        const interval = setInterval(() => {
            fetchLatest()
        }, 3000)
        return () => clearInterval(interval)
    })

    return (<div >
        <label>{prop.kind}: </label>
            <span style={{backgroundColor: `${color}`, margin:"5px", padding: '0px 10px', width: "10px", fontWeight:"bolder"}}>
                {temp} &deg;C
            </span>
        </div>)
}
