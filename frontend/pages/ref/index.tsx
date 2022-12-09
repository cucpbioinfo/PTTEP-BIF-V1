// import { PageLayout } from 'components/new/PageLayout'
// import { PrivateRoute } from 'components/PrivateRoute'
// import { Checkbox } from 'antd'
// import React from 'react'
// import type { CheckboxChangeEvent } from 'antd/es/checkbox';
// import { useRouter } from 'next/router'
// //import { SunburstChart } from 'features/sunburstd3/SunbrustChart'
// const index = () => {
//   const router = useRouter()
//   const { locale } = router

//   const onFilterChange = (key: string, value?: string) => {
//     const url = {
//       pathname: router.pathname,
//       query: {
//         ...router.query,
//         [key]: value,
//       },
//     }
//     router.push(url, url, { locale })
//   }

  
//   const onChange = (e: CheckboxChangeEvent) => {
//     console.log(`checked = ${e.target.checked}`);
//       (value: string | undefined) => {
//         onFilterChange('checkvalue', value)
//       }
//   };
  
//   return (
//     <PageLayout>
//       {/* <h2>Sunburst Chart by d3 comp</h2>
//       <SunburstChart/> */}
//         <Checkbox onChange={onChange}>Checkbox</Checkbox>

//     </PageLayout>
//   )
// }

// export default PrivateRoute(index)


// import classNames from 'classnames';
// import RcCheckbox from 'rc-checkbox';
// import * as React from 'react';
// import { useContext } from 'react';
// import { ConfigContext } from '../config-provider';
// import { FormItemInputContext } from '../form/context';
// import warning from '../_util/warning';
// import { GroupContext } from './Group';
// import DisabledContext from '../config-provider/DisabledContext';

// export interface AbstractCheckboxProps<T> {
//   prefixCls?: string;
//   className?: string;
//   defaultChecked?: boolean;
//   checked?: boolean;
//   style?: React.CSSProperties;
//   disabled?: boolean;
//   onChange?: (e: T) => void;
//   onClick?: React.MouseEventHandler<HTMLElement>;
//   onMouseEnter?: React.MouseEventHandler<HTMLElement>;
//   onMouseLeave?: React.MouseEventHandler<HTMLElement>;
//   onKeyPress?: React.KeyboardEventHandler<HTMLElement>;
//   onKeyDown?: React.KeyboardEventHandler<HTMLElement>;
//   value?: any;
//   tabIndex?: number;
//   name?: string;
//   children?: React.ReactNode;
//   id?: string;
//   autoFocus?: boolean;
//   type?: string;
//   skipGroup?: boolean;
// }

// export interface CheckboxChangeEventTarget extends CheckboxProps {
//   checked: boolean;
// }

// export interface CheckboxChangeEvent {
//   target: CheckboxChangeEventTarget;
//   stopPropagation: () => void;
//   preventDefault: () => void;
//   nativeEvent: MouseEvent;
// }

// export interface CheckboxProps extends AbstractCheckboxProps<CheckboxChangeEvent> {
//   indeterminate?: boolean;
// }

// const InternalCheckbox: React.ForwardRefRenderFunction<HTMLInputElement, CheckboxProps> = (
//   {
//     prefixCls: customizePrefixCls,
//     className,
//     children,
//     indeterminate = false,
//     style,
//     onMouseEnter,
//     onMouseLeave,
//     skipGroup = false,
//     disabled,
//     ...restProps
//   },
//   ref,
// ) => {
//   const { getPrefixCls, direction } = React.useContext(ConfigContext);
//   const checkboxGroup = React.useContext(GroupContext);
//   const { isFormItemInput } = useContext(FormItemInputContext);
//   const contextDisabled = useContext(DisabledContext);
//   const mergedDisabled = disabled || checkboxGroup?.disabled || contextDisabled;

//   const prevValue = React.useRef(restProps.value);

//   React.useEffect(() => {
//     checkboxGroup?.registerValue(restProps.value);
//     warning(
//       'checked' in restProps || !!checkboxGroup || !('value' in restProps),
//       'Checkbox',
//       '`value` is not a valid prop, do you mean `checked`?',
//     );
//   }, []);

//   React.useEffect(() => {
//     if (skipGroup) {
//       return;
//     }
//     if (restProps.value !== prevValue.current) {
//       checkboxGroup?.cancelValue(prevValue.current);
//       checkboxGroup?.registerValue(restProps.value);
//       prevValue.current = restProps.value;
//     }
//     return () => checkboxGroup?.cancelValue(restProps.value);
//   }, [restProps.value]);

//   const prefixCls = getPrefixCls('checkbox', customizePrefixCls);
//   const checkboxProps: CheckboxProps = { ...restProps };
//   if (checkboxGroup && !skipGroup) {
//     checkboxProps.onChange = (...args) => {
//       if (restProps.onChange) {
//         restProps.onChange(...args);
//       }
//       if (checkboxGroup.toggleOption) {
//         checkboxGroup.toggleOption({ label: children, value: restProps.value });
//       }
//     };
//     checkboxProps.name = checkboxGroup.name;
//     checkboxProps.checked = checkboxGroup.value.indexOf(restProps.value) !== -1;
//   }
//   const classString = classNames(
//     {
//       [`${prefixCls}-wrapper`]: true,
//       [`${prefixCls}-rtl`]: direction === 'rtl',
//       [`${prefixCls}-wrapper-checked`]: checkboxProps.checked,
//       [`${prefixCls}-wrapper-disabled`]: mergedDisabled,
//       [`${prefixCls}-wrapper-in-form-item`]: isFormItemInput,
//     },
//     className,
//   );
//   const checkboxClass = classNames({
//     [`${prefixCls}-indeterminate`]: indeterminate,
//   });
//   const ariaChecked = indeterminate ? 'mixed' : undefined;
//   return (
//     // eslint-disable-next-line jsx-a11y/label-has-associated-control
//     <label
//       className={classString}
//       style={style}
//       onMouseEnter={onMouseEnter}
//       onMouseLeave={onMouseLeave}
//     >
//       <RcCheckbox
//         aria-checked={ariaChecked}
//         {...checkboxProps}
//         prefixCls={prefixCls}
//         className={checkboxClass}
//         disabled={mergedDisabled}
//         ref={ref}
//       />
//       {children !== undefined && <span>{children}</span>}
//     </label>
//   );
// };

// const Checkbox = React.forwardRef<unknown, CheckboxProps>(InternalCheckbox);
// if (process.env.NODE_ENV !== 'production') {
//   Checkbox.displayName = 'Checkbox';
// }

// export default Checkbox;

import { Button, Checkbox, Form,Divider,Tooltip} from "antd";
import * as React from "react";
import { CheckBoxTest } from 'features/test/checktest'
import { AntdCheckBoxTest }  from 'features/test/antdCheck'
import { DefEvenness,DefDiversity,DefNumber,DefSam,Diversityinfo } from "components/new/DefInfo";
import { InfoCircleOutlined } from '@ant-design/icons';


// interface MyFormValues {
//   permissions: {
//     orderChargeCards: boolean;
//     orderChargePoint: boolean;
//   };
// }

export default function App() {
  // const [form] = Form.useForm<MyFormValues>();

  // const initialValues = {
  //   permissions: {
  //     orderChargeCards: false,
  //     orderChargePoint: false
  //   }
  // };

  return (
    <>
    <DefDiversity/>
    <Divider/>
    <DefEvenness/>
    <Divider/>
    <DefNumber/>
    <Divider/>
    <b>Sampling levels</b>
    <DefSam/>
    <Divider/>
    <div className="flex w-full items-center">
      <div className="flex w-full items-center divide-x mr-4">
        <div>
          <div className="mr-4 ml-2"></div>
        </div>
      </div>
      <div>
        <div className="mr-2">
          <Tooltip placement="top" title="def"><InfoCircleOutlined/></Tooltip>
        </div>
      </div>
    </div>
    <Diversityinfo/>
    {/* <div className="App">
      <Form form={form} onFinish={console.log} initialValues={initialValues}>
        <Form.Item
          name={["permissions", "orderChargePoint"]}
          valuePropName="checked"
        >
          <Checkbox>Beställa laddboxar</Checkbox>
        </Form.Item>

        <Form.Item
          name={["permissions", "orderChargeCards"]}
          valuePropName="checked"
        >
          <Checkbox>Beställa laddbrickor</Checkbox>
        </Form.Item>

        <Button type="primary" htmlType="submit">
          Submit
        </Button>
      </Form>
      <CheckBoxTest/>
      <AntdCheckBoxTest/>
    </div> */}
    </>
  );
}