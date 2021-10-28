import { Dispatch, SetStateAction } from 'react';
import styles from './code-input.module.css';

interface Props {
  value: string;
  setValue: Dispatch<SetStateAction<string>>;
}

const CodeInput = (props: Props) => {
  return (
    <textarea
      className={styles.input}
      onChange={(e) => props.setValue(e.target.value)}
    >
      {props.value}
    </textarea>
  );
};

export default CodeInput;
