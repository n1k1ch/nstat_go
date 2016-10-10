import * as React from "react";
import EntrySchema from "./EntrySchema";

export class EntryRow extends React.Component<{entry:EntrySchema}, {}> {
    render() {
        return (
            <tr>
                <td>{this.props.entry.Id}</td>
                <td>{this.props.entry.Description}</td>
                <td>{this.props.entry.Amount}</td>
                <td>{this.props.entry.Currency}</td>
            </tr>
        );
    }
}