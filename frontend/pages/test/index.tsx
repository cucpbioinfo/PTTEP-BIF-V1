import { Button, Checkbox, Form,Divider,Tooltip} from "antd";
import * as React from "react";
import { PageLayout } from 'components/new/PageLayout'
import { CheckBoxTest } from 'features/test/checktest'
import { AntdCheckBoxTest }  from 'features/test/antdCheck'
import { DefEvenness,DefDiversity,DefNumber,DefSam,Diversityinfo } from "components/new/DefInfo";
import { InfoCircleOutlined } from '@ant-design/icons';

export default function App() {

  return (
   <>
   <PageLayout>
   {/* <div className="h-auto w-3/4 m-auto border-1">
aa
   </div> */}
    <ul role="list" className="divide-y divide-slate-200 list-none hover:list-disc ">
      {/* ------------- */}
      <li className="flex h-auto w-3/4 m-auto border-1 box-border items-center justify-between">
        <div className="flex box-border items-start w-full m-auto p-auto">
          {/* <div className="m-auto p-auto box-border">
            <span className="m-auto p-auto box-border relative inline-block overflow-hidden whitespace-nowrap text-center align-middle outline-none h-24 w-24 rounded-full outline-none">
              <img src="https://randomuser.me/api/portraits/men/87.jpg"/> 
            </span>
          </div> */}
          <div className="m-auto p-auto w-full box-border">
            <h4 className="m-auto p-auto box-border outline-none">
              <a href="https://ant.design">Louis</a>
            </h4>
            <div className="m-auto p-auto box-border outline-none">celestin.louis@example.com</div>
          </div>
        </div>
        {/* <div>Content</div> */}
      </li>
      {/* ------------- */}
            
    </ul>
   </PageLayout>
   </>
  );
}
