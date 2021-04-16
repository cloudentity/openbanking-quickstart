import React from 'react';
import {Bar, BarChart, CartesianGrid, Cell, ResponsiveContainer, XAxis, YAxis} from 'recharts';

export default function AnalyticsBarChart({data, filtering, onChangeFiltering}) {

  return (
    <ResponsiveContainer width={'100%'} height={300}>
      <BarChart
        data={data}
        margin={{
          top: 5, right: 30, left: 20, bottom: 5,
        }}
      >
        <CartesianGrid strokeDasharray="3 3"/>
        <XAxis dataKey="name" axisLine={false} tickLine={false}/>
        <YAxis axisLine={false} tickLine={false}/>
        <Bar
          dataKey="value"
          background={{fill: '#eee'}}
          onClick={e => !filtering?.months.includes(e.name)
            ? onChangeFiltering({months: [e.name], categories: []})
            : onChangeFiltering({months: [], categories: []})
          }
        >
          {data.map((entry, index) => (
            <Cell cursor="pointer" fill={filtering?.months.includes(entry.name) ? '#36C6AF' : '#1F2D48'} key={`cell-${index}`}/>
          ))}
        </Bar>
      </BarChart>
    </ResponsiveContainer>
  )
}
