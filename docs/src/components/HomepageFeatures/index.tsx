import type {ReactNode} from 'react';
import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';

type FeatureItem = {
    title: string;
    Svg: React.ComponentType<React.ComponentProps<'svg'>>;
    description: ReactNode;
};

const FeatureList: FeatureItem[] = [
    {
        title: 'Resource management automation',
        Svg: require('@site/static/img/art_resource_management.svg').default,
        description: (
            <>
                The operator allows you to automate the processes of deployment, updating and managing resources
                Casdoor. This simplifies the processing of configurations for organizations, applications and
                permissions, reducing the likelihood errors and saving developers' time.
            </>
        ),
    },
    {
        title: 'Scalability and flexibility',
        Svg: require('@site/static/img/art_scalability.svg').default,
        description: (
            <>
                You can quickly add or delete resources for various organizations and applications, as well as manage
                Permissions on the fly, which makes the system more adaptive to changes.
            </>
        ),
    },
    {
        title: 'Simplified observation and state management',
        Svg: require('@site/static/img/art_observation.svg').default,
        description: (
            <>
                The operator ensures the monitoring of the Casdoor resource state and automatically responds to
                changes, which allows maintaining high availability and reliability. This is especially important for
                accounting for organizations and applications, where it is necessary to monitor the permissions and
                access in real time.
            </>
        ),
    },
];

function Feature({title, Svg, description}: FeatureItem) {
    return (
        <div className={clsx('col col--4')}>
            <div className="text--center">
                <Svg className={styles.featureSvg} role="img"/>
            </div>
            <div className="text--center padding-horiz--md">
                <Heading as="h3">{title}</Heading>
                <p>{description}</p>
            </div>
        </div>
    );
}

export default function HomepageFeatures(): ReactNode {
    return (
        <section className={styles.features}>
            <div className="container">
                <div className="row">
                    {FeatureList.map((props, idx) => (
                        <Feature key={idx} {...props} />
                    ))}
                </div>
            </div>
        </section>
    );
}
