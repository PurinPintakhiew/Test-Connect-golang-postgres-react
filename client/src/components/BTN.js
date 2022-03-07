import {useState} from 'react';

const BTN = () => {
    const [data,setData] = useState([]);

    const getData = () => {
        fetch('http://localhost:4444/lists')
        .then((res) => res.json())
        .then((json) => setData(json))
    }

    return (
        <div>
            <button onClick={getData}>BTN</button>
            <table>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                    </tr>
                </thead>
                <tbody>
            {data.map((val) => {
                return (
                    <tr>
                        <td>{val.id}</td>
                        <td>{val.name}</td>
                    </tr>
                )
            })}
            </tbody>
            </table>
        </div>
    )
}

export default BTN;