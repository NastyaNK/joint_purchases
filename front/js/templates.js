let loaded_templates = [];

replace_template = (text, key, value) => {
    return text.replaceAll("{% " + key + " %}", value == null ? '' : value);
};

template = async(name, row = null) => {
    let file = "/mvp/templates/" + name + ".html";
    let text;
    if (file in loaded_templates) {
        text = loaded_templates[file];
    } else {
        text = await send(file);
        text = text[1];
        loaded_templates[file] = text;
    }
    if (row != null)
        Object.keys(row).forEach((key) => text = replace_template(text, key, row[key]));
    return text;
};

multi_template = async(name, result, where = null, else_ = null, default_else = '') => {
    let file = "/mvp/templates/" + name + ".html";
    let temp;
    if (file in loaded_templates) {
        temp = loaded_templates[file];
    } else {
        temp = await send(file);
        temp = temp[1];
        loaded_templates[file] = temp;
    }
    let output = "";
    if (result != null && result.length > 0) {
        result.forEach((row) => {
            let output_part = temp;
            Object.keys(row).forEach((key) => output_part = replace_template(output_part, key, row[key]));
            if (where != null) {
                Object.keys(where).forEach((key) => {
                    if (key in row) {
                        Object.keys(where[key]).forEach((key2) => {
                            if (row[key] in where[key][key2]) {
                                output_part = replace_template(output_part, key2, where[key][key2][row[key]]);
                            } else if (else_ != null && key2 in else_) {
                                output_part = replace_template(output_part, key2, else_[key2])
                            } else {
                                output_part = replace_template(output_part, key2, default_else)
                            }
                        });
                    }
                });
            }
            output += output_part;
        });
    }
    return output;
};