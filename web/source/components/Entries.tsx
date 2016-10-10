import * as React from "react";
import EntrySchema from "./EntrySchema";
import {EntryRow} from "./EntryRow";

export class Entries extends React.Component<{}, {entries: EntrySchema[]}> {
    constructor() {
        super();
        var arr:EntrySchema[] = [];
        arr.push({Id: 1, Description: "Hie"});
        arr.push({Id: 2});
        this.state = {entries: arr};
    }

    render() {
        return (
            <table>
                <thead>
                    <tr>
                        <th>Id</th>
                        <th>Description</th>
                        <th>Amount</th>
                        <th>Currency</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        this.state.entries.map((val) => {
                            return (
                                <EntryRow key={val.Id}
                                          entry={val}
                                />
                                )
                            })
                        }
                </tbody>
            </table>
        );
    }

    renderEntries(entries:EntrySchema[]) {
        return (
            entries.map((element) => {

            })
        );
    }
}