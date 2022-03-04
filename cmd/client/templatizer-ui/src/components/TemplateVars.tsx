import { Controller } from "react-hook-form";
import TextField from "@material-ui/core/TextField";
import { Grid } from "@material-ui/core";
import { Var } from './Client';

interface TemplateVarsProps {
    control: any;
    setValue?: any;
    updateParam: any;
}

export const TemplateVars: React.FC<TemplateVarsProps> = ({
    control,
    setValue,
    updateParam,
}) => {

    return (


        <Controller
            name="templateVars"
            control={control}
            render={({ field: { value, onChange } }) => {

                const templateVarsFields = value.variables.map((variable: Var) => {

                    const name = `var_${variable.name}`
                    setValue(name, variable.default)
                    return (
                        <Grid item lg={10}>
                            <Controller
                                name={name}
                                control={control}
                                render={({
                                    field: { value },
                                }) => (
                                    <TextField
                                        fullWidth
                                        value={value}
                                        label={variable.name}
                                        helperText={variable.description}
                                        variant="standard"
                                        onChange={(val: any) => {
                                            setValue(name, val.target.value)
                                            updateParam(variable.name, val.target.value)
                                        }}
                                    />
                                )}
                            />
                        </Grid>
                    )

                });

                return (
                    <Grid container spacing={2}>
                        {templateVarsFields}
                    </Grid>
                )
            }}
        />
    );
};
