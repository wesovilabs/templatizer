import { Controller, useFormContext } from "react-hook-form";
import TextField from "@material-ui/core/TextField";
import { FormControl, FormControlLabel, FormLabel, Grid, InputLabel, MenuItem, Radio, RadioGroup, Select, Typography } from "@material-ui/core";
import { useState } from "react";

interface RepositoryDetailsProps {
    control: any;
    setValue?: any;
}

export const defaultValues = {
    url: "https://github.com/ivancorrales/seed.git",
    branchDefault: "default",
    authMechanism: 'none',
    configPath: '.templatizer.yml'
};




export const RepositoryDetails: React.FC<RepositoryDetailsProps> = ({
    control,
    setValue,
}) => {
    const [showCustomBranch, setCustomBranch] = useState<boolean>(false);
    const [showBasicAuth, setBasicAuth] = useState<boolean>(false);
    const [showTokenAuth, setTokenAuth] = useState<boolean>(false);


    return (
        <Grid container spacing={1}>
            <Grid item lg={12}>
                <Typography variant="h6">Repository Details</Typography>
            </Grid>

            <Grid item lg={8}>
                <Controller name="url" control={control} render={({
                    field: { onChange, value },
                    fieldState: { error },
                    formState,
                }) => (
                    <TextField
                        fullWidth
                        label="URL"
                        helperText={"i.e https://github.com/ivancorrales/seed.git"}
                        error={!!error}
                        onChange={onChange}
                        value={value}
                        variant="standard"

                    />
                )} />
            </Grid>
            <Grid item lg={12}>
                <FormLabel id="branch-type" component="legend">Branch</FormLabel>
            </Grid>
            <Grid item lg={3}>
                <Controller
                    name={"branchDefault"}
                    control={control}
                    render={({
                        field: { value },
                    }) => (
                        <RadioGroup
                            aria-labelledby="branch-type"
                            value={value} onChange={(val) => {
                                setCustomBranch(val.target.value != 'default');
                                setValue('branchDefault', val.target.value);
                            }}>
                            <FormControlLabel
                                value="default"
                                label="Default branch"
                                control={<Radio />}
                            />
                            <FormControlLabel
                                value="no-default"
                                label="Other branch"
                                control={<Radio />}
                            />
                        </RadioGroup>

                    )}
                />
            </Grid>
            {showCustomBranch && (
                <Grid item lg={9}>
                    <Controller name="branch" control={control} render={({
                        field: { onChange, value },
                        fieldState: { error },
                    }) => (
                        <TextField
                            fullWidth
                            label="Name"
                            helperText={"develop, stable, release/v1.0.1"}
                            error={!!error}
                            onChange={onChange}
                            value={value}
                            variant="standard"
                        />
                    )} />

                </Grid>
            )}
            <Grid item lg={12}>
                <Typography variant="h6">Authentication mechanism</Typography>
            </Grid>
            <Grid item lg={3}>
                <Controller
                    name={"authMechanism"}
                    control={control}
                    render={({
                        field: { onChange, value },
                        fieldState: { error },
                        formState,
                    }) => (
                        <RadioGroup
                            aria-labelledby="branch-type"
                            name="branchDefault"
                            value={value}
                            onChange={(val) => {
                                setTokenAuth(val.target.value == 'token');
                                setBasicAuth(val.target.value == 'basic');
                                setValue('authMechanism', val.target.value);
                            }}>
                            <FormControlLabel
                                value="none"
                                label="None"
                                control={<Radio />}
                            />
                            <FormControlLabel
                                value="basic"
                                label="Basic"
                                control={<Radio />}
                            />
                            <FormControlLabel
                                value="token"
                                label="Token"
                                control={<Radio />}
                            />
                        </RadioGroup>
                    )}
                />
            </Grid>

            <Grid container spacing={2}>

                <Grid item lg={5} hidden={!showBasicAuth}>
                    <Controller name="authUsername" control={control} render={({
                        field: { onChange, value },
                        fieldState: { error },
                        formState,
                    }) => (
                        <TextField
                            fullWidth
                            label="Username"
                            helperText={"develop, stable, release/v1.0.1"}
                            error={!!error}
                            onChange={onChange}
                            value={value}
                            variant="standard"
                        />
                    )} />
                </Grid>
                <Grid item lg={5} hidden={!showBasicAuth}>
                    <Controller name="authPassword" control={control} render={({
                        field: { onChange, value },
                        fieldState: { error },
                        formState,
                    }) => (
                        <TextField
                            fullWidth
                            label="Password"
                            helperText={"develop, stable, release/v1.0.1"}
                            error={!!error}
                            onChange={onChange}
                            value={value}
                            variant="standard"
                        />
                    )} />

                </Grid>

                <Grid item lg={10} hidden={!showTokenAuth}>
                    <Controller name="authToken" control={control} render={({
                        field: { onChange, value },
                        fieldState: { error },
                        formState,
                    }) => (
                        <TextField
                            fullWidth
                            label="Token"
                            helperText={"develop, stable, release/v1.0.1"}
                            error={!!error}
                            onChange={onChange}
                            value={value}
                            variant="standard"
                        />
                    )} />
                </Grid>
            </Grid>
            <Grid item lg={12}>
                <Typography variant="h6">Configuration</Typography>
            </Grid>
            <Grid item lg={8}>
                <Controller name="configPath" control={control} render={({
                    field: { onChange, value },
                    fieldState: { error },
                    formState,
                }) => (
                    <TextField
                        fullWidth
                        label="Path"
                        helperText={"i.e .template/templatizer.yml"}
                        error={!!error}
                        onChange={onChange}
                        value={value}
                        variant="standard"

                    />
                )} />
            </Grid>
        </Grid >
    );
};
